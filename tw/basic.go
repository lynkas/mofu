package tw

import (
	"github.com/g8rswimmer/go-twitter"
)

type tweet struct {
	twitter.TweetObj
}

type Author struct {
	Id          string `json:"id"`
	DisplayName string `json:"name"`
	UserName    string `json:"username"`
}

func (a *Author) Username() string {
	return a.UserName
}

func (a *Author) Name() string {
	return a.DisplayName
}

func (a *Author) ToAuthor() *Author {
	return a
}

func (a *Author) ID() string {
	return a.Id
}

func (a *Author) From(obj *twitter.UserObj) *Author {
	a.DisplayName = obj.Name
	a.UserName = obj.UserName
	a.Id = obj.ID
	return a
}

type Tweet struct {
	Id string `json:"id"`
}

func (t *Tweet) ID() string {
	return t.Id
}

func (t *Tweet) ToTweet() *Tweet {
	return t
}

func (t *Tweet) From(obj *twitter.TweetObj) *Tweet {
	t.Id = obj.ID
	return t
}

type Media struct {
	Key  string `json:"media_key"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (m *Media) From(obj *twitter.MediaObj) *Media {
	m.URL = obj.URL
	m.Type = obj.Type
	m.Key = obj.Key
	return m
}

type AuthorTweet struct {
	Author *Author
	Tweet  *Tweet
}

func (t *tweet) ID() string {
	return t.TweetObj.ID
}

func (t *tweet) ToTweet() *Tweet {
	return &Tweet{Id: t.ID()}
}

type tweetAuthor struct {
	twitter.UserObj
}

func (t *tweetAuthor) Username() string {
	return t.UserName
}

func (t *tweetAuthor) Name() string {
	return t.UserObj.Name
}

func (t *tweetAuthor) ID() string {
	return t.UserObj.ID
}

func (t *tweetAuthor) ToAuthor() *Author {
	return &Author{
		Id:          t.ID(),
		DisplayName: t.Name(),
		UserName:    t.Username(),
	}
}

type ID interface {
	ID() string
}

type ITweet interface {
	ID
	ToTweet() *Tweet
}

type IAuthor interface {
	Username() string
	Name() string
	ToAuthor() *Author
	ID
}
type IMedia interface {
	Key() string
	Type() string
	URL() string
}

type ICompoundMedia interface {
	IMedia
	Tweet() ITweet
	Author() IAuthor
	Broadcast() IAuthor
}

type IManyMedias interface {
	Media() []ICompoundMedia
}

type TwitterQueryResult struct {
	timeline      *twitter.UserTimeline
	previousMaxID string
}

func (t *TwitterQueryResult) Media() []ICompoundMedia {
	if t.timeline.Includes == nil {
		return nil
	}
	result := make([]ICompoundMedia, len(t.timeline.Includes.Medias))
	for i, media := range t.timeline.Includes.Medias {
		result[i] = &UserTimelineMedia{
			TwitterQueryResult: t,
			mediaKey:           media.Key,
		}
	}
	return result
}

func (t *TwitterQueryResult) MaxID() string {
	if t.timeline.Meta.NewestID == "" {
		return t.previousMaxID
	}
	return t.timeline.Meta.NewestID
}
