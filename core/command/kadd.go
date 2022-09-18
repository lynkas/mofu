package command

import (
	"errors"
	"mofu/value"
	"regexp"
)

type kadd struct {
	command
	username string
	key      string
}

func (c *kadd) parse() error {
	c.trim()
	if c.empty() {
		return errors.New("should provide usernames")
	}
	space := regexp.MustCompile("\\s+")
	split := space.Split(c.args, -1)
	if len(split) < 2 {
		return errors.New("not valid")
	}
	c.username, c.key = split[0], split[1]
	return nil
}

func (c *kadd) Run() (value.MessageMakeup, error) {
	_, err := c.commandRunner.AddSubscriptions([]string{c.username})
	if err != nil {
		return nil, err
	}
	msg, err := c.commandRunner.GetHistory(c.key)
	if err != nil {
		return nil, err
	}
	return value.ToEditMediaMessage(msg), nil
}
