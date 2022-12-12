package value

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type text struct {
	content string
}

func NewText(content string) *text {
	return &text{content: content}
}

func (t *text) ImageURL() string {
	return ""
}

func (t *text) Content() string {
	return t.content
}

func (t *text) Keyboard() *tgbotapi.InlineKeyboardMarkup {
	return nil
}

func (t *text) Message(channelID int64, replyID int) tgbotapi.Chattable {
	return TextMessage(channelID, replyID, t)
}

func (t *text) EditMessage(channelID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(channelID, replyID, t)
}

func TextMessage(channelID int64, replyID int, t IMessage) tgbotapi.Chattable {
	msg := tgbotapi.NewMessage(channelID, t.Content())
	msg.ReplyToMessageID = replyID
	msg.ReplyMarkup = t.Keyboard()
	msg.ParseMode = tgbotapi.ModeHTML
	return msg
}

//func ToSendTextMessage(t IMessage) MessageMakeup {
//	return func(chatID int64, replyID int) tgbotapi.Chattable {
//		return TextMessage(chatID, replyID, t)
//	}
//}
