package value

import (
	"mofu/tw"
)

type singleControlMessage struct {
	*controlMessage
	isSent bool
}

func NewSingleControlMessage(media tw.ICompoundMedia, authorNotFollowed bool, isSent bool) *singleControlMessage {
	return &singleControlMessage{
		NewControlMessage(media, authorNotFollowed),
		isSent,
	}
}

func (c *singleControlMessage) Content() string {
	return BuildContent(c)
}

func (c *singleControlMessage) AdditionalContent() string {
	content := c.controlMessage.AdditionalContent()
	if c.isSent {
		content += "\n 这张图发过了"
	}
	return content
}
