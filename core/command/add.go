package command

import (
	"errors"
	"mofu/value"
	"regexp"
	"strings"
)

type add struct {
	command
	usernames []string
}

func (c *add) parse() error {
	c.trim()
	if c.empty() {
		return errors.New("should provide usernames")
	}
	space := regexp.MustCompile("\\s+")
	c.usernames = space.Split(c.args, -1)
	for i, username := range c.usernames {
		c.usernames[i] = extractUsername(username)
	}
	return nil
}

func extractUsername(content string) string {
	content = strings.ReplaceAll(content, " ", "")
	content = strings.Replace(content, "@", "", 1)

	linkRegex := regexp.MustCompile("https://twitter.com/([0-9a-zA-Z_]{1,15}).*?")
	subMatch := linkRegex.FindStringSubmatch(content)
	if len(subMatch) >= 2 {
		return subMatch[1]
	}
	return content
}

func (c *add) Run() (value.MessageMakeup, error) {
	msg, err := c.commandRunner.AddSubscriptions(c.usernames)
	if err != nil {
		return nil, err
	}
	return value.ToSendTextMessage(msg), err
}
