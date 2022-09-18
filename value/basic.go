package value

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mofu/tw"
	"strings"
)

type IMessageMakeup interface {
	Message(chatID int64, replyID int) tgbotapi.Chattable
	EditMessage(chatID int64, replyID int) tgbotapi.Chattable
}

type MessageMakeup func(chatID int64, replyID int) tgbotapi.Chattable

type IMessage interface {
	ImageURL() string
	Content() string
	Keyboard() *tgbotapi.InlineKeyboardMarkup
	IMessageMakeup
}

type MessageFactory struct {
	Media               tw.ICompoundMedia
	Nsfw                int
	Approved            int
	Decided             int
	AuthorIsNotFollowed bool
	Operator            string
}

func (m *MessageFactory) ToControlMessage() *controlMessage {
	return &controlMessage{
		MediaMessage:        MediaMessage{media: m.Media},
		AuthorIsNotFollowed: m.AuthorIsNotFollowed,
	}

}

//
//// TODO edited message
//
//func (m *MessageFactory) ToMessage() IMessage {
//	if m.Controlled == Controlled {
//		if m.Decided == Decided {
//			if m.IsSending {
//				return m.ToPublicMessage()
//			} else {
//				return m.ToProcessedControlMessage()
//
//			}
//		} else {
//
//		}
//	}
//	return m.ToControlMessage()
//
//}

func (m *MessageFactory) ToPublicMessage() *publicMessage {
	return &publicMessage{
		MediaMessage: MediaMessage{
			media: m.Media,
		},
		nsfw: m.Nsfw == Nsfw,
	}

}

func (m *MessageFactory) ToProcessedControlMessage() *ProcessedControlMessage {
	return &ProcessedControlMessage{
		controlMessage: *m.ToControlMessage(),
		decided:        m.Decided,
		approved:       m.Approved,
		nsfw:           m.Nsfw,
		operator:       m.Operator,
	}
}

type ContentBuildable interface {
	Author() string
	Tags() []string
	AdditionalContent() string
}

func BuildContent(c ContentBuildable) string {
	result := strings.Builder{}
	result.WriteString(c.Author())

	if c.Tags() != nil {
		result.WriteString("\n")
	}
	for _, tag := range c.Tags() {
		result.WriteString(fmt.Sprintf("#%s ", tag))
	}

	if c.AdditionalContent() != "" {
		result.WriteString("\n" + c.AdditionalContent())
	}
	return result.String()
}
