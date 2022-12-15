package core

import (
	"context"
	log "github.com/sirupsen/logrus"
	"mofu/ent"
	"mofu/ent/history"
	"mofu/ent/subscription"
	"mofu/tw"
	"mofu/utils"
	"mofu/value"
	"time"
)

type IMessageUpdater interface {
	Can() bool
	Next() ([]*ent.History, error)
	After(histories *ent.History) error
	MakeMessage(h *ent.History) value.IMessage
}

type MessageUpdater struct {
	*utils.PeriodFrequency
	db *ent.Client
}

func NewMessageUpdater(count int, gap time.Duration, db *ent.Client) *MessageUpdater {
	return &MessageUpdater{
		PeriodFrequency: utils.NewFrequencyLimit(count, gap),
		db:              db,
	}
}

func NewSendUpdater(msgUpdater *MessageUpdater, cooldown time.Duration, together int, enoughOrNothing bool) *SendUpdater {
	return &SendUpdater{
		MessageUpdater:  msgUpdater,
		cooldown:        cooldown,
		together:        together,
		enoughOrNothing: enoughOrNothing,
	}
}

type SendUpdater struct {
	*MessageUpdater
	cooldown        time.Duration
	together        int
	enoughOrNothing bool
}

func (s *SendUpdater) Next() ([]*ent.History, error) {
	result, err := s.query()
	if s.enoughOrNothing && len(result) == s.together || !s.enoughOrNothing {
		return result, err
	}
	return nil, err
}

func (s *SendUpdater) query() ([]*ent.History, error) {
	result, err := s.db.History.Query().Where(
		history.And(
			history.SentFlagEQ(value.Controlled|value.Decided|value.Approved|value.No),
			history.LastUpdateLT(time.Now().Add(-s.cooldown)),
			history.SendingContentNEQ([]byte("{}"))),
	).Limit(s.together).All(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return result, err
}

func (s *SendUpdater) After(history *ent.History) error {
	if history == nil {
		return nil
	}
	err := s.db.History.UpdateOne(history).SetSentFlag(history.SentFlag | value.Sent).Exec(context.Background())
	return err
}

func (s *SendUpdater) MakeMessage(history *ent.History) value.IMessage {
	return value.NewPublicMessage(tw.SingleMediaResultFrom(history.SendingContent), history.ContentFlag)
}

type DecideUpdater struct {
	*MessageUpdater
}

func NewDecideUpdater(msgUpdater *MessageUpdater) *DecideUpdater {
	return &DecideUpdater{
		MessageUpdater: msgUpdater,
	}
}

func (s *DecideUpdater) Next() ([]*ent.History, error) {
	result, err := s.query()
	if result == nil {
		return nil, err
	}
	return []*ent.History{result}, err
}

func (s *DecideUpdater) query() (*ent.History, error) {
	result, err := s.db.History.Query().Where(
		history.And(
			history.Or(
				history.SentFlagLT(value.Controlled),
				history.SentFlagIsNil(),
			),
			history.SendingContentNEQ([]byte("{}")),
			history.MentionedCountGTE(2),
		),
	).
		Order(ent.Asc(history.FieldLastUpdate)).
		First(context.Background())
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return result, err
}

func (s *DecideUpdater) After(history *ent.History) error {
	if history == nil {
		return nil
	}
	err := s.db.History.UpdateOne(history).SetSentFlag(value.Controlled).Exec(context.Background())
	return err
}

func (s *DecideUpdater) authorFollowed(userID string) bool {
	result, err := s.db.Subscription.Query().Where(
		subscription.ID(userID)).
		Exist(context.Background())
	if err != nil {
		log.Warn(err)
	}
	return result
}

func (s *DecideUpdater) MakeMessage(history *ent.History) value.IMessage {
	media := tw.SingleMediaResultFrom(history.SendingContent)
	return value.NewControlMessage(media, !s.authorFollowed(media.Author().ID()))
}
