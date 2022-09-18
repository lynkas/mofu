package command

import (
	"errors"
	"mofu/value"
	"strconv"
	"strings"
)

type decide struct {
	command
	id       string
	sending  int
	content  int
	operator string
}

func flag(command string) (int, int, error) {
	flags := strings.SplitN(command, ".", 2)
	flagNums := []int{-1, -1}

	for i, flag := range flags {
		num, err := strconv.Atoi(flag)
		if err != nil {
			return -1, -1, err
		}
		flagNums[i] = num
	}
	return flagNums[0], flagNums[1], nil
}

func (c *decide) parse() error {
	split := strings.SplitN(c.args, "_", 2)
	if len(split) != 2 {
		return errors.New("format between _ not correct")
	}
	var err error
	c.sending, c.content, err = flag(split[0])
	if err != nil {
		return err
	}
	c.id = split[1]
	return nil
}

func (c *decide) Run() (value.MessageMakeup, error) {
	msg, err := c.commandRunner.UpdateHistoryFlag(c.id, c.sending, c.content, c.operator)
	if err != nil {
		return nil, err
	}
	return value.ToEditMediaMessage(msg), err
}
