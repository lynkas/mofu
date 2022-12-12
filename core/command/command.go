package command

import (
	"mofu/value"
	"strings"
)

type IRunner interface {
	AddSubscriptions(usernames []string) (value.IMessage, error)
	RemoveSubscription(username string) (value.IMessage, error)
	UpdateHistoryFlag(key string, send int, content int, operator string) (value.IMessage, error)
	GetHistory(key string) (value.IMessage, error)
	WebAuth(user string) (value.IMessage, error)
	WebDestroy(token string) (value.IMessage, error)
	ListSettings() (value.IMessage, error)
	SetSetting(key, val string) (value.IMessage, error)
	RemoveSetting(key string) (value.IMessage, error)
}

type ICommand interface {
	parse() error
	Run() (value.MessageMakeup, error)
}

type command struct {
	commandRunner IRunner
	function      string
	args          string
}

func (c *command) parse() error { return nil }
func (c *command) empty() bool {
	return c.args == ""
}
func (c *command) trim() {
	c.args = strings.Trim(c.args, " ")
}

func (c *command) Run() (value.MessageMakeup, error) {
	return nil, nil
}

const (
	Decide     = "/decide"
	Log        = "/log"
	Setting    = "/setting"
	Add        = "/add"
	Kadd       = "/kadd"
	Remove     = "/remove"
	WebAuth    = "/web_auth"
	WebDestroy = "/web_destroy"
)

func New(function, args, operator string, runner IRunner) (ICommand, error) {
	c := command{
		commandRunner: runner,
		function:      function,
		args:          args,
	}
	var commandObj ICommand
	switch function {
	case Decide:
		commandObj = &decide{command: c, operator: operator}
	case Add:
		commandObj = &add{command: c}
	case Kadd:
		commandObj = &kadd{command: c}
	case Log:
		commandObj = &testLog{command: c}
	case Setting:
		commandObj = &setting{command: c}
	case Remove:
		commandObj = &remove{command: c}
	case WebAuth:
		commandObj = &webAuth{command: c, user: operator}
	case WebDestroy:
		commandObj = &webDestroy{command: c}
	default:
		commandObj = &c
	}

	err := commandObj.parse()
	if err != nil {
		return nil, err
	}
	return commandObj, nil
}
