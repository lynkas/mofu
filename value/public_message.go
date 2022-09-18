package value

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"mofu/tw"
)

type publicMessage struct {
	MediaMessage
	nsfw bool
}

func NewPublicMessage(media tw.ICompoundMedia, nsfwFlag int) *publicMessage {
	return &publicMessage{
		MediaMessage: MediaMessage{media: media},
		nsfw:         Nsfw&nsfwFlag != 0,
	}
}

func (p *publicMessage) Tags() []string {
	var tags []string
	if p.nsfw {
		tags = append(tags, "nsfw")
	}
	tags = append(tags, p.MediaMessage.media.Author().Username())
	return tags
}

func (p *publicMessage) Content() string {
	return BuildContent(p)
}

func (p *publicMessage) Message(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, p)
}

func (p *publicMessage) EditMessage(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, p)
}
