package command

import (
	"mofu/value"
	"strings"
)

type setting struct {
	command
	item  *string
	value *string
}

func (c *setting) parse() error {
	split := strings.SplitN(c.args, " ", 2)
	if len(split) >= 1 {
		c.item = &split[0]
	}
	if len(split) >= 2 {
		c.value = &split[1]
	}
	return nil
}

func (c *setting) Run() (result value.MessageMakeup, err error) {
	var msg value.IMessage
	if c.item == nil {
		msg, err = c.commandRunner.ListSettings()
	} else if c.value == nil {
		msg, err = c.commandRunner.RemoveSetting(*c.item)
	} else {
		msg, err = c.commandRunner.SetSetting(*c.item, *c.value)
		if err != nil {
			return nil, err
		}
	}

	return value.ToSendTextMessage(msg), err
}
