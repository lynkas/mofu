package core

import (
	"mofu/core/command"
	"mofu/value"
)

func (c *Core) CommandProcess(function, args, operator string) (value.MessageMakeup, error) {
	commander, err := command.New(function, args, operator, c)
	if err != nil {
		return value.ToSendMediaMessage(value.NewText(err.Error())), err
	}
	return commander.Run()
}

func (c *Core) SingleLink(tweetID string) ([]value.MessageMakeup, error) {
	query, err := c.twitter.QueryTweetsByID([]string{tweetID})
	if err != nil {
		return nil, err
	}
	var result []value.MessageMakeup
	medias, _, err := c.ResultProcess(query, true)
	if err != nil {
		return nil, err
	}

	for _, media := range medias {
		h, _ := c.getHistoryByID(media.Key())
		result = append(result, value.ToSendMediaMessage(
			value.NewSingleControlMessage(media, !c.HasSubscription(media.Author().ID()), h.SentFlag&value.Sent != value.No)))
	}

	return result, nil
}
