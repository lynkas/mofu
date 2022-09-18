package tw

import (
	"github.com/g8rswimmer/go-twitter"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
)

type UserTimelineMedia struct {
	*TwitterQueryResult
	mediaKey string
}

func (u *UserTimelineMedia) Key() string {
	return u.mediaKey
}

func (u *UserTimelineMedia) mediaObj() *twitter.MediaObj {
	for _, media := range u.timeline.Includes.Medias {
		if media.Key == u.Key() {
			return &media
		}
	}
	return nil
}

func (u *UserTimelineMedia) Type() string {
	return u.mediaObj().Type
}

func (u *UserTimelineMedia) URL() string {
	return u.mediaObj().URL
}

func (u *UserTimelineMedia) Tweet() ITweet {
	return u.getTweetByMediaKey(u.mediaKey)
}

func (u *UserTimelineMedia) getTweetByMediaKey(key string) *tweet {
	for _, t := range append(u.timeline.Includes.Tweets, u.timeline.Tweets...) {
		for _, mediaKey := range t.Attachments.MediaKeys {
			if mediaKey == key {
				return &tweet{t}
			}
		}
	}
	return nil
}

func (u *UserTimelineMedia) Author() IAuthor {
	t := u.getTweetByMediaKey(u.mediaKey)
	if t.ReferencedTweets == nil {
		return u.getUser(t.AuthorID)
	}
	//if len(t.ReferencedTweets) == 1 && t.ReferencedTweets[0].Type == "retweeted" {
	//
	//	return u.getUser(t.AuthorID)
	//}
	if slices.Contains(t.Attachments.MediaKeys, u.mediaKey) {
		return u.Broadcast()
	}

	log.Warn("take care of this, Author", t)
	return nil
}

func (u *UserTimelineMedia) getUser(id string) *tweetAuthor {
	for _, user := range u.timeline.Includes.Users {
		if user.ID == id {
			return &tweetAuthor{user}
		}
	}
	return nil
}

func (u *UserTimelineMedia) Broadcast() IAuthor {
	return u.getUser(u.timeline.Tweets[0].AuthorID)
}

func Format(u ICompoundMedia) SingleMediaResult {
	return SingleMediaResult{
		Original: &AuthorTweet{
			Author: u.Author().ToAuthor(),
			Tweet:  u.Tweet().ToTweet(),
		},
		BroadCast: u.Broadcast().ToAuthor(),
		Media: &Media{
			Key:  u.Key(),
			Type: u.Type(),
			URL:  u.URL(),
		},
	}
}
