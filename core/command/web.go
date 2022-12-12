package command

import (
	"mofu/value"
)

type webAuth struct {
	command
	user string
}

func (c *webAuth) Run() (value.MessageMakeup, error) {
	msg, err := c.commandRunner.WebAuth(c.user)
	if err != nil {
		return nil, err
	}
	return value.ToTextMessage(msg), err
}

type webDestroy struct {
	command
	token string
}

func (c *webDestroy) Run() (value.MessageMakeup, error) {
	msg, err := c.commandRunner.WebDestroy(c.token)
	if err != nil {
		return nil, err
	}
	return value.ToTextMessage(msg), err
}
