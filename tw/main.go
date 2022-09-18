package tw

import (
	"context"
	"fmt"
	"github.com/g8rswimmer/go-twitter"
	log "github.com/sirupsen/logrus"
	"mofu/utils"
	"net/http"
	"time"
)

func UserTimelineOpts() twitter.UserTimelineOpts {
	return twitter.UserTimelineOpts{
		TweetFields: []twitter.TweetField{
			twitter.TweetFieldAttachments,
			twitter.TweetFieldAuthorID,
			twitter.TweetFieldText,
		},
		MediaFields: []twitter.MediaField{
			twitter.MediaFieldURL,
			twitter.MediaFieldPreviewImageURL,
		},
		UserFields: []twitter.UserField{
			twitter.UserFieldName,
			twitter.UserFieldUserName,
			twitter.UserFieldURL,
		},
		Expansions: []twitter.Expansion{
			twitter.ExpansionAttachmentsMediaKeys,
			twitter.ExpansionAuthorID,
			twitter.ExpansionReferencedTweetsIDAuthorID,
		},
	}
}
func TweetOpts() twitter.TweetFieldOptions {
	return twitter.TweetFieldOptions{
		TweetFields: []twitter.TweetField{
			twitter.TweetFieldAttachments,
			twitter.TweetFieldAuthorID,
			twitter.TweetFieldReferencedTweets,
			twitter.TweetFieldText,
		},
		MediaFields: []twitter.MediaField{
			twitter.MediaFieldURL,
			twitter.MediaFieldPreviewImageURL,
			twitter.MediaFieldMediaKey,
		},
		UserFields: []twitter.UserField{
			twitter.UserFieldName,
			twitter.UserFieldUserName,
			twitter.UserFieldURL,
		},
		Expansions: []twitter.Expansion{
			twitter.ExpansionAttachmentsMediaKeys,
			twitter.ExpansionAuthorID,
			twitter.ExpansionReferencedTweetsID,
		},
	}
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

type TwitterClient struct {
	twitterUserLoader *twitter.User
	tweetLoader       *twitter.Tweet
	frequencyLimit    *utils.PeriodFrequency
}

func New(TwitterApiToken string) *TwitterClient {
	twitterUserLoader := &twitter.User{
		Authorizer: authorize{
			Token: TwitterApiToken,
		},
		Client: &http.Client{},
		Host:   "https://api.twitter.com",
	}

	tweetLoader := &twitter.Tweet{
		Authorizer: authorize{
			Token: TwitterApiToken,
		},
		Client: &http.Client{},
		Host:   "https://api.twitter.com",
	}
	limit := utils.NewFrequencyLimit(880, time.Minute*15)
	t := &TwitterClient{
		twitterUserLoader: twitterUserLoader,
		tweetLoader:       tweetLoader,
		frequencyLimit:    limit,
	}
	return t
}

func (t *TwitterClient) CanGet() {
	t.frequencyLimit.Can()
}

func (t *TwitterClient) UserLookupByUsername(usernames []string) ([]*twitter.UserObj, error) {
	t.CanGet()
	fieldOpts := twitter.UserFieldOptions{}
	lookups, err := t.twitterUserLoader.LookupUsername(context.Background(), usernames, fieldOpts)
	var result []*twitter.UserObj
	if err != nil {
		return nil, err
	}
	for i, lookup := range lookups {
		for _, username := range usernames {
			if lookup.User.UserName == username {
				user := lookups[i].User
				result = append(result, &user)
			}
		}
	}

	return result, nil
}

func (t *TwitterClient) QueryTweetsWithUserID(userID, lastID string, maxCount int) (*TwitterQueryResult, error) {
	t.CanGet()
	tweetOpts := UserTimelineOpts()
	tweetOpts.SinceID = lastID
	tweetOpts.MaxResults = maxCount
	userTweets, err := t.twitterUserLoader.Tweets(context.Background(), userID, tweetOpts)

	if err != nil {
		log.Warning(err)
		return nil, err
	}

	return &TwitterQueryResult{
		timeline:      userTweets,
		previousMaxID: lastID,
	}, nil
}

func (t *TwitterClient) GetUpdateLandmark(userID string) (string, error) {
	update, err := t.QueryTweetsWithUserID(userID, "", 5)
	if err != nil {
		return "", err
	}
	return update.MaxID(), err
}

func (t *TwitterClient) QueryTweetsByID(tweetIDs []string) (*IDQueryResult, error) {
	tweetOpts := TweetOpts()
	tweets, err := t.tweetLoader.Lookup(context.Background(), tweetIDs, tweetOpts)

	if err != nil {
		return nil, err
	}

	return &IDQueryResult{lookup: &tweets}, nil
}
