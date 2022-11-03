package value

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mofu/tw"
	"strings"
)

type controlMessage struct {
	MediaMessage
	AuthorIsNotFollowed bool
}

func NewControlMessage(media tw.ICompoundMedia, authorNotFollowed bool) *controlMessage {
	return &controlMessage{
		MediaMessage:        MediaMessage{media: media},
		AuthorIsNotFollowed: authorNotFollowed,
	}
}

func (c *controlMessage) Author() string {
	if c.isRetweet() {
		return c.MediaMessage.Author() + fmt.Sprintf(` | <a href="https://twitter.com/%s">%s</a>`,
			c.media.Broadcast().Username(), c.media.Broadcast().Name())
	} else {
		return c.MediaMessage.Author()
	}
}

func (c *controlMessage) Tags() []string {
	return []string{"tw" + c.media.Key()}
}

func (c *controlMessage) AdditionalContent() string {
	builder := strings.Builder{}
	builder.WriteString(c.media.URL())
	if c.AuthorIsNotFollowed {
		builder.WriteString(" 未关注的作者")
	}
	return builder.String()
}

func (c *controlMessage) isRetweet() bool {
	return c.media.Author().ID() != c.media.Broadcast().ID()
}

func (c *controlMessage) Content() string {
	return BuildContent(c)
}

func (c *controlMessage) Message(chatID int64, replyID int) tgbotapi.Chattable {
	return Message(chatID, replyID, c)
}
func (c *controlMessage) EditMessage(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, c)
}

func (c *controlMessage) Keyboard() *tgbotapi.InlineKeyboardMarkup {
	first := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("发！", fmt.Sprintf("/decide %d.%d_%s", Controlled|Decided|Approved, No, c.media.Key())),
		tgbotapi.NewInlineKeyboardButtonData("NSFW，发！", fmt.Sprintf("/decide %d.%d_%s", Controlled|Decided|Approved, Nsfw, c.media.Key())),
	)
	if c.AuthorIsNotFollowed {
		first = append(first, tgbotapi.NewInlineKeyboardButtonData("关注原作者", fmt.Sprintf("/kadd %s %s", c.media.Author().Username(), c.media.Key())))
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(first)
	return &keyboard
}
