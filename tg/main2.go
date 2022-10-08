package tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	olog "log"
	"mofu/value"
	"regexp"
	"strings"
	"time"
)

type ITelegramFunc interface {
	CommandProcess(command string, args string, operator string) (value.MessageMakeup, error)
	SingleLink(tweetID string) ([]value.MessageMakeup, error)
	DecideTask() chan value.MessageMakeup
	SendTask() chan value.MessageMakeup
}

func (t *Telegram) command(commandFull, operator string, messageID int) {
	chatID := t.controlRoomID

	if !strings.HasPrefix(commandFull, "/") {
		commandFull = "/decide " + commandFull
	}
	commands := strings.SplitN(commandFull, " ", 2)
	if commands[0] == "/log" {
		chatID = t.warningRoomID
	}
	var msg value.MessageMakeup
	if len(commands) != 2 {
		commands = append(commands, "")
	}
	var err error
	msg, err = t.function.CommandProcess(commands[0], commands[1], operator)
	if err != nil {
		msg = value.ToSendTextMessage(value.NewText(err.Error()))
	}
	if msg == nil {
		return
	}
	_, err = t.Request(msg(chatID, messageID))
	if err != nil {
		log.Warning(err)
		return
	}
}

func (t *Telegram) chatCommand(update tgbotapi.Update) {
	t.command(update.Message.Text, fullName(update), update.Message.MessageID)

}
func (t *Telegram) keyboardCommand(update tgbotapi.Update) {
	t.command(update.CallbackQuery.Data, fullName(update), update.CallbackQuery.Message.MessageID)
}

func extractTweetID(content string) string {
	linkRegex := regexp.MustCompile("https://twitter.com/[0-9a-zA-Z_]{1,15}/status/(\\d+.*?)")
	subMatch := linkRegex.FindStringSubmatch(content)
	if len(subMatch) >= 2 {
		return subMatch[1]
	}
	return content
}

func validTweetID(content string) bool {
	idRegex := regexp.MustCompile("(\\d+)")
	return extractTweetID(content) != content || idRegex.MatchString(content)
}

func (t *Telegram) processSingle(update tgbotapi.Update) {
	id := extractTweetID(update.Message.Text)
	medias, err := t.function.SingleLink(id)
	if err != nil {
		log.Warning(err)
	}
	for _, msg := range medias {
		_, err = t.Send(msg(update.Message.Chat.ID, update.Message.MessageID))
		if err != nil {
			log.Warning(err)
		}
	}
}

func (t *Telegram) RunTask() {
	var msg value.MessageMakeup
	for {
		select {
		case msg = <-t.function.DecideTask():
			t.Send(msg(t.controlRoomID, 0))
		case msg = <-t.function.SendTask():
			t.Send(msg(t.broadcastRoomID, 0))
			//default :
			//	time.Sleep(time.Second * 2)
		}
	}

}

func (t *Telegram) Update() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := t.bot.GetUpdatesChan(updateConfig)
	for {
		for update := range updates {
			if update.CallbackQuery != nil {
				t.keyboardCommand(update)
			} else if update.Message != nil {
				if update.Message.IsCommand() {
					if t.validCommand(update) {
						t.chatCommand(update)
					}
				} else {
					if validTweetID(update.Message.Text) {
						t.processSingle(update)
					} else {
						if !update.Message.From.IsBot {
						}
					}
				}

			} else {
				//TODO
			}
		}
		olog.Println("bot failed")
		time.Sleep(time.Minute)
	}
}

func fullName(update tgbotapi.Update) string {
	return fmt.Sprintf("%s %s", update.SentFrom().FirstName, update.SentFrom().LastName)
}
