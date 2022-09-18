package value

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ProcessedControlMessage struct {
	controlMessage
	decided  int
	approved int
	nsfw     int
	operator string
}

func (c *ProcessedControlMessage) Content() string {
	return BuildContent(c)
}

func (c *ProcessedControlMessage) Tags() []string {
	tags := []string{c.media.Key()}
	if c.approved == Approved {
		tags = append(tags, "通过")
	} else {
		tags = append(tags, "未通过")
	}

	if c.nsfw == Nsfw {
		tags = append(tags, "nsfw")
	}

	return tags
}

func (c *ProcessedControlMessage) AdditionalContent() string {
	return fmt.Sprintf(">> %s", c.operator)
}

func (c *ProcessedControlMessage) Message(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, c)
}

func (c *ProcessedControlMessage) EditMessage(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, c)
}

func EditMessage(chatID int64, replyID int, c IMessage) tgbotapi.Chattable {
	edit := tgbotapi.NewEditMessageCaption(chatID, replyID, c.Content())
	edit.ParseMode = tgbotapi.ModeHTML
	edit.ReplyMarkup = c.Keyboard()
	return edit
}

func (c *ProcessedControlMessage) Keyboard() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("撤销选项", fmt.Sprintf("/decide %d.%d_%s", Controlled, No, c.media.Key())),
		),
	)
	return &keyboard
}
