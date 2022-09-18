package tw

import "github.com/g8rswimmer/go-twitter"

type IDQueryMedia struct {
	tweet      *twitter.TweetLookup
	mediaIndex int
}

type IDQueryResult struct {
	lookup *twitter.TweetLookups
}

func (i *IDQueryMedia) mediaObj() *twitter.MediaObj {
	return i.tweet.AttachmentMedia[i.mediaIndex]
}

func (i *IDQueryMedia) Key() string {
	return i.mediaObj().Key
}

func (i *IDQueryMedia) Type() string {
	return i.mediaObj().Type
}

func (i *IDQueryMedia) URL() string {
	return i.mediaObj().URL

}

func (i *IDQueryMedia) Author() IAuthor {
	return &tweetAuthor{*i.tweet.User}
}

func (i *IDQueryMedia) Tweet() ITweet {
	return &tweet{i.tweet.Tweet}
}

func (i *IDQueryMedia) Broadcast() IAuthor {
	return i.Author()
}

func (i *IDQueryResult) Media() []ICompoundMedia {
	var result []ICompoundMedia
	for _, t := range *i.lookup {
		for k := range t.AttachmentMedia {
			result = append(result, &IDQueryMedia{
				tweet:      &t,
				mediaIndex: k,
			})
		}
	}
	return result
}
