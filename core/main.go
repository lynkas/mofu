package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/g8rswimmer/go-twitter"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
	"mofu/ent"
	"mofu/ent/auth"
	"mofu/ent/author"
	"mofu/ent/history"
	"mofu/ent/subscription"
	"mofu/tw"
	"mofu/utils"
	"mofu/value"
	"time"
)

type Core struct {
	twitter *tw.TwitterClient
	db      *ent.Client
	decide  chan value.MessageMakeup
	send    chan value.MessageMakeup
	setting *Setting
}

func New(db *ent.Client, twitter *tw.TwitterClient) *Core {
	settings := NewSetting(db.Setting)
	c := &Core{
		twitter: twitter,
		db:      db,
		decide:  make(chan value.MessageMakeup),
		send:    make(chan value.MessageMakeup),
		setting: settings,
	}
	return c
}

func (c *Core) ListSettings() (value.IMessage, error) {
	result := value.NewOperationResults("settings")
	keys := maps.Keys(Defaults)
	slices.Sort(keys)
	for _, k := range keys {
		dft := Defaults[k]
		if v, ok := c.setting.cache[k]; ok {
			result.Set(k, v)
		} else {
			result.Set(k, fmt.Sprintf("%s/>d", dft))
		}
	}

	return result, nil
}

func (c *Core) Auth(ctx context.Context, token string) *string {
	a, err := c.db.Auth.Query().Where(auth.Token(token)).Only(ctx)
	if err != nil {
		return nil
	}
	return &a.User
}

func (c *Core) ListHistory(ctx context.Context, timeBefore *int64, offset, limit int) ([]*ent.History, error) {
	q := c.db.History.Query()
	if timeBefore != nil {
		q = q.Where(history.TakeEffectTimeLTE(time.Unix(*timeBefore/1000, 0)))
	}
	return q.Where(history.MentionedCountGTE(2)).Offset(offset).Limit(limit).Order(ent.Desc(history.FieldTakeEffectTime), ent.Desc(history.FieldCreateAt)).All(ctx)
}

func (c *Core) AddSubscriptions(usernames []string) (value.IMessage, error) {

	users, err := c.twitter.UserLookupByUsername(usernames)
	if err != nil {
		return nil, err
	}
	result := value.NewOperationResults("添加")
	for _, username := range usernames {
		result.Set(username, "未找到")
	}
	c.addSubscription(users, result)

	return result, nil
}

func (c *Core) addSubscription(users []*twitter.UserObj, result value.IMapResult) value.IMapResult {
	for _, user := range users {
		_, err := c.db.Subscription.Create().
			SetID(user.ID).
			SetUsername(user.UserName).
			SetName(user.Name).Save(context.Background())
		if err != nil {
			if errors.Is(err, ent.ConstraintError{}) {
				result.Set(user.UserName, "加过了")
			} else {
				result.Set(user.UserName, err.Error())
			}
		} else {
			result.Set(user.UserName, "已添加")
		}
	}
	return result
}

func (c *Core) removeSubscription(username string) error {
	_, err := c.db.Subscription.Delete().Where(subscription.UsernameEQ(username)).Exec(context.Background())
	return err
}

func (c *Core) RemoveSubscription(username string) (value.IMessage, error) {
	if username == "" {
		return nil, errors.New("should provide a username")
	}
	var result = value.NewOperationResults("删除")
	result.Set(username, "成功")
	err := c.removeSubscription(username)
	if err != nil {
		result.Set(username, err.Error())
	}
	return result, err
}

func (c *Core) getHistoryByID(mediaKey string) (*ent.History, error) {
	h, err := c.db.History.Get(context.Background(), mediaKey)
	return h, err
}

func (c *Core) addHistory(media tw.ICompoundMedia, asSentControlled bool) error {
	sentFlag := value.No
	if asSentControlled {
		sentFlag = value.Controlled
	}
	err := c.db.History.Create().
		SetID(media.Key()).
		SetCreatorID(media.Author().ID()).
		SetSentFlag(sentFlag).
		SetSendingContent(tw.Format(media).Dump()).
		Exec(context.Background())
	return err
}

func (c *Core) addMentionedCount(key string) error {
	h, err := c.db.History.UpdateOneID(key).AddMentionedCount(1).Save(context.Background())
	if err != nil {
		return err
	}
	if h.MentionedCount == 2 {
		err = h.Update().SetTakeEffectTime(time.Now()).Exec(context.Background())
		if err != nil {
			return err
		}
	}
	return err
}

func (c *Core) hasHistory(mediaKey string) (bool, error) {
	return c.db.History.Query().Where(history.ID(mediaKey)).Exist(context.Background())
}

//
//func (c *Core) updateHistoryStatus(id string, controlled, decided, approved, sent, nsfw int) (*ent.History, error) {
//	data, err := c.db.History.UpdateOneID(id).SetSentFlag(controlled | decided | approved | sent).SetContentFlag(nsfw).Save(context.Background())
//	return data, err
//}

func (c *Core) SetSetting(key, val string) (value.IMessage, error) {
	err := c.setting.Set(key, val)
	if err != nil {
		return nil, err
	}
	return c.ListSettings()
}
func (c *Core) RemoveSetting(key string) (value.IMessage, error) {
	err := c.setting.Remove(key)
	if err != nil {
		return nil, err
	}
	return c.ListSettings()
}

func (c *Core) sent(id string) (bool, error) {
	return c.db.History.
		Query().
		Where(
			history.ID(id),
			history.SentFlagGTE(value.Controlled|value.Decided|value.Approved|value.Sent),
		).
		Exist(context.Background())
}
func (c *Core) UpdateHistoryFlag(id string, sent, nsfw int, operator string) (value.IMessage, error) {
	data, err := c.getHistoryByID(id)
	if err != nil {
		return nil, err
	}
	if data.SentFlag&value.Approved != 0 && sent&value.Approved != 0 {
		sent = data.SentFlag
	} else {
		data, err = c.db.History.UpdateOneID(id).SetSentFlag(sent).SetContentFlag(nsfw).Save(context.Background())
		if err != nil {
			return nil, err
		}
	}

	media := tw.SingleMediaResultFrom(data.SendingContent)
	msg := &value.MessageFactory{
		AuthorIsNotFollowed: !c.HasSubscription(media.Author().ID()),
		Media:               media,
		Nsfw:                nsfw & value.Nsfw,
		Approved:            sent & value.Approved,
		Decided:             sent & value.Decided,
		Operator:            operator,
	}
	if msg.Decided == value.No {
		return msg.ToControlMessage(), nil
	}
	return msg.ToProcessedControlMessage(), nil
}

func (c *Core) nextShouldUpdateSubscription(gap time.Duration) (*ent.Subscription, error) {
	result, err := c.db.Subscription.Query().Order(ent.Asc(subscription.FieldLastUpdate)).First(context.Background())
	if ent.IsNotFound(err) {
		return result, nil
	}
	if err != nil {
		return result, err
	}

	if time.Now().Sub(result.LastUpdate) > gap {
		err = c.db.Subscription.UpdateOneID(result.ID).SetLastUpdate(time.Now()).Exec(context.Background())
		if err != nil {
			return nil, err
		}
		return result, err
	}

	return nil, err
}

func (c *Core) Update() {
	go c.updateSend()
	go c.updateControl()
}

func (c *Core) UpdateSubscribe() {
	for {
		updated := false
		updated = updated || c.updateSubscribe()
		<-time.After(time.Second * time.Duration(2+rand.Intn(10)/10.0))
	}
}

func (c *Core) updateSubscribe() bool {
	sub, err := c.nextShouldUpdateSubscription(c.setting.GetDuration(UpdateGap, time.Minute*15))
	if sub != nil {
		log.Debug("update ", sub.Name, "\t", sub.Username)
	} else {
		log.Debug("no update")
	}
	if err != nil {
		log.Warn(err)
		return false
	}
	if sub == nil {
		return false
	}
	medias, err := c.GetUserNewTweets(sub.ID, sub.LastLandmark, 20)
	if err != nil {
		log.Warn(fmt.Sprintf("Update failed: \n"+
			"%s\n\n"+
			"<code>%s</code> 更新失败", err.Error(), sub.Username))
		return false
	}

	return len(medias) != 0
}

func (c *Core) updateControl() {
	updater := NewMessageUpdater(10, time.Second, c)
	for {
		updater := NewDecideUpdater(updater)
		updated := UpdaterFunc(updater, c.decide)
		if !updated {
			<-time.After(time.Second * 2)
		}
	}
}
func (c *Core) updateSend() {
	for {
		//TODO
		prevSendGap := c.setting.GetDuration(SendingGap, time.Minute*15)
		updater := NewMessageUpdater(1, prevSendGap, c)
		for {
			updater := NewSendUpdater(updater,
				c.setting.GetDuration(WaitBeforeDecided, time.Minute*15),
				c.setting.GetInt(NumberSendingTogether, 5),
				c.setting.GetBool(EnoughOrNoting, true))
			updated := UpdaterFunc(updater, c.send)
			if !updated {
				<-time.After(time.Second * 2)
			}
			if prevSendGap != c.setting.GetDuration(SendingGap, time.Minute*15) {
				break
			}
		}
	}

}

func UpdaterFunc(updater IMessageUpdater, msgChan chan value.MessageMakeup) bool {
	updater.Can()
	histories, err := updater.Next()
	if err != nil {
		log.Warn(err)
	}
	var msgs = make([]value.MessageMakeup, len(histories))
	for i, h := range histories {
		msgs[i] = value.ToSendMediaMessage(updater.MakeMessage(h))
	}
	for i, msg := range msgs {
		msgChan <- msg
		errs := updater.After(histories[i])
		if errs != nil {
			log.Warn(errs)
		}
	}
	return len(msgs) != 0

}

func (c *Core) DecideTask() chan value.MessageMakeup {
	return c.decide
}
func (c *Core) SendTask() chan value.MessageMakeup {
	return c.send
}

func (c *Core) SetLandmark(userID, landmark string) bool {
	err := c.db.Subscription.UpdateOneID(userID).SetLastLandmark(landmark).Exec(context.Background())
	if err != nil {
		log.Warn(err)
		return false
	}
	return err == nil
}

func (c *Core) GetUserNewTweets(userID, lastID string, maxCount int) ([]tw.ICompoundMedia, error) {
	query, err := c.twitter.QueryTweetsWithUserID(userID, lastID, maxCount)
	if err != nil {
		return nil, err
	}
	newLandMark := query.MaxID()
	c.SetLandmark(userID, newLandMark)
	medias, _, err := c.ResultProcess(query, false)
	return medias, err
	//var result []value.IMessage
	//for i, exist := range exists {
	//	if exist {
	//		continue
	//	}
	//	result = append(result, (&value.MessageFactory{
	//		Media: medias[i],
	//	}).ToControlMessage())
	//}

}

func (c *Core) legacyUpdate(media tw.ICompoundMedia) {
	h, err := c.getHistoryByID(media.Key())
	if err != nil {
		log.Warn(err)
	}
	if !bytes.Equal(h.SendingContent, []byte("{}")) {
		return
	}
	err = c.db.History.UpdateOneID(h.ID).
		SetCreatorID(media.Author().ID()).
		SetSentFlag(value.Controlled).
		SetSendingContent(tw.Format(media).Dump()).
		Exec(context.Background())
	if err != nil {
		log.Warn(err)
	}

}

func (c *Core) ResultProcess(query tw.IManyMedias, asSentControlled bool) ([]tw.ICompoundMedia, []bool, error) {
	medias := make([]tw.ICompoundMedia, len(query.Media()))
	exist := make([]bool, len(medias))
	for i, media := range query.Media() {
		have, err := c.hasHistory(media.Key())
		exist[i] = have
		if err != nil {
			return nil, nil, err
		}
		if !have {
			err = c.addHistory(media, asSentControlled)
			if err != nil {
				return nil, nil, err
			}
		} else {
			c.legacyUpdate(media)
			err := c.addMentionedCount(media.Key())
			if err != nil {
				return nil, nil, err
			}
		}
		medias[i] = media
	}
	return medias, exist, nil
}

func (c *Core) HasSubscription(userID string) bool {
	result, err := c.db.Subscription.Query().Where(
		subscription.ID(userID)).
		Exist(context.Background())
	if err != nil {
		log.Warn(err)
	}
	return result
}

func (c *Core) GetHistory(key string) (value.IMessage, error) {
	h, err := c.getHistoryByID(key)
	if err != nil {
		return nil, err
	}
	media := tw.SingleMediaResultFrom(h.SendingContent)
	return value.NewControlMessage(media, !c.HasSubscription(media.Author().ID())), nil
}
func (c *Core) WebAuth(name string) (value.IMessage, error) {
	randomString, err := utils.GenerateRandomString(128)
	if err != nil {
		return nil, err
	}
	_, err = c.db.Auth.Create().SetToken(randomString).SetUser(name).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return value.NewText(fmt.Sprintf("<code>%s</code>", randomString)), nil
}
func (c *Core) WebDestroy(key string) (value.IMessage, error) {
	_, err := c.db.Auth.Delete().Where(auth.Token(key)).Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return value.NewText("删了"), nil

}
func (c *Core) QueryAuthor(ctx context.Context, name string) ([]*ent.Author, error) {
	return c.db.Author.Query().Where(author.UserNameContains(name)).Limit(10).Order(ent.Asc(author.FieldUserName)).All(ctx)
}

func (c *Core) QueryAuthorMedia(ctx context.Context, id string, limit, offset int) ([]*ent.History, error) {
	return c.db.History.Query().Where(history.CreatorID(id)).Limit(limit).Offset(offset).All(ctx)
}

func (c *Core) addAuthor(id, name string) error {
	ctx := context.Background()
	author, err := c.db.Author.Query().Where(author.UserID(id)).Only(ctx)
	if ent.IsNotFound(err) {
		creator := c.db.Author.Create()
		authorUpdater(creator.Mutation(), id, name)
		_, err = creator.Save(ctx)
	} else {
		updater := author.Update()
		authorUpdater(updater.Mutation(), id, name)
		_, err = updater.Save(ctx)
	}
	return err
}

func authorUpdater(a *ent.AuthorMutation, id, name string) {
	a.SetUserID(id)
	a.SetUserName(name)
}
