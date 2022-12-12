package value

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mofu/tw"
)

type MediaMessage struct {
	media tw.ICompoundMedia
}

func (c *MediaMessage) ImageURL() string {
	return c.media.URL()
}

func (c *MediaMessage) Content() string {
	return BuildContent(c)
}

func (c *MediaMessage) Keyboard() *tgbotapi.InlineKeyboardMarkup {
	return nil
}

func (c *MediaMessage) Author() string {
	return fmt.Sprintf(`<a href="https://twitter.com/%s/status/%s">%s</a>`,
		c.media.Author().Username(), c.media.Tweet().ID(), c.media.Author().Name())
}

func (c *MediaMessage) Tags() []string {
	return nil
}

func (c *MediaMessage) AdditionalContent() string {
	return ""
}
func (c *MediaMessage) Message(chatID int64, replyID int) tgbotapi.Chattable {
	return Message(chatID, replyID, c)
}
func (c *MediaMessage) EditMessage(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, c)
}

func Message(chatID int64, replyID int, c IMessage) tgbotapi.Chattable {
	config := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(c.ImageURL()))
	config.ParseMode = tgbotapi.ModeHTML
	config.Caption = c.Content()
	config.ReplyToMessageID = replyID
	config.ReplyMarkup = c.Keyboard()
	return config
}

type MessageMakeup interface {
	ToChattable(chatID int64, replyID int) tgbotapi.Chattable
	ToMsg() IMessage
}

type msg struct {
	IMessage
}

func (m *msg) ToMsg() IMessage {
	return m
}

type editMessage struct {
	msg
}
type sendMessage struct {
	msg
}
type textMessage struct {
	msg
}

func (c *editMessage) ToChattable(chatID int64, replyID int) tgbotapi.Chattable {
	edit := tgbotapi.NewEditMessageCaption(chatID, replyID, c.Content())
	edit.ParseMode = tgbotapi.ModeHTML
	edit.ReplyMarkup = c.Keyboard()
	return edit
}

func (c *sendMessage) ToChattable(chatID int64, replyID int) tgbotapi.Chattable {
	config := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(c.ImageURL()))
	config.ParseMode = tgbotapi.ModeHTML
	config.Caption = c.Content()
	config.ReplyToMessageID = replyID
	config.ReplyMarkup = c.Keyboard()
	return config
}
func (c *textMessage) ToChattable(chatID int64, replyID int) tgbotapi.Chattable {
	return TextMessage(chatID, replyID, c)
}

func ToTextMessage(m IMessage) *textMessage {
	return &textMessage{msg{m}}
}
func ToEditMediaMessage(m IMessage) *editMessage {
	return &editMessage{msg{m}}
}
func ToSendMediaMessage(m IMessage) *sendMessage {
	return &sendMessage{msg{m}}
}

/*
func ToSendMediaMessage(c IMessage) MessageMakeup {
	return func(chatID int64, replyID int) tgbotapi.Chattable {
		config := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(c.ImageURL()))
		config.ParseMode = tgbotapi.ModeHTML
		config.Caption = c.Content()
		config.ReplyToMessageID = replyID
		config.ReplyMarkup = c.Keyboard()
		return config
	}
}

func ToEditMediaMessage(c IMessage) MessageMakeup {
	return func(chatID int64, replyID int) tgbotapi.Chattable {
		edit := tgbotapi.NewEditMessageCaption(chatID, replyID, c.Content())
		edit.ParseMode = tgbotapi.ModeHTML
		edit.ReplyMarkup = c.Keyboard()
		return edit
	}
}*/
