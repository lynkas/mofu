package command

import (
	"mofu/value"
)

type remove struct {
	command
	cleanedName string
}

func (c *remove) parse() error {
	c.trim()
	if !c.empty() {
		c.cleanedName = c.args
	}
	return nil
}

func (c *remove) Run() (value.MessageMakeup, error) {
	msg, err := c.commandRunner.RemoveSubscription(c.cleanedName)
	if err != nil {
		return nil, err
	}
	return value.ToSendTextMessage(msg), err
}
