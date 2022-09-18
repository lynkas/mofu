package command

import (
	log "github.com/sirupsen/logrus"
	"mofu/value"
)

type testLog struct {
	command
}

func (c *testLog) parse() error {
	return nil
}

func (c *testLog) Run() (value.MessageMakeup, error) {
	log.Warn(c.args)
	return nil, nil
}
