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

func (t *Telegram) command(commandFull string, update tgbotapi.Update, messageID int) {
	chatID := t.controlRoomID
	operator := update.SentFrom()
	if !strings.HasPrefix(commandFull, "/") {
		commandFull = "/decide " + commandFull
	}
	commands := strings.SplitN(commandFull, " ", 2)
	if commands[0] == "/log" {
		chatID = t.warningRoomID
	}

	if strings.HasPrefix(commandFull, "/web") {
		if !t.isAdmin(operator) {
			log.Warning("/web without auth")
			log.Warning(operator)
			return
		}
		if !update.FromChat().IsPrivate() {
			msg := value.TextMessage(update.FromChat().ID, update.Message.MessageID, value.NewText("私聊"))
			_, err := t.Request(msg)
			if err != nil {
				log.Warning(err)
				return
			}
		}
	}

	var msg value.MessageMakeup
	if len(commands) != 2 {
		commands = append(commands, "")
	}
	var err error
	msg, err = t.function.CommandProcess(commands[0], commands[1], fullName(operator))
	if err != nil {
		msg = value.ToTextMessage(value.NewText(err.Error()))
	}
	if msg == nil {
		return
	}
	_, err = t.Request(msg.ToChattable(chatID, messageID))
	if err != nil {
		log.Warning(err)
		return
	}
}

func (t *Telegram) chatCommand(update tgbotapi.Update) {
	t.command(update.Message.Text, update, update.Message.MessageID)

}
func (t *Telegram) keyboardCommand(update tgbotapi.Update) {
	t.command(update.CallbackQuery.Data, update, update.CallbackQuery.Message.MessageID)
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
		_, err = t.Send(msg.ToChattable(update.Message.Chat.ID, update.Message.MessageID))
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
			t.Send(msg.ToChattable(t.controlRoomID, 0))
		case msg = <-t.function.SendTask():
			t.Send(msg.ToChattable(t.broadcastRoomID, 0))
			//default :
			//	time.Sleep(time.Second * 2)
		}
	}

}

func (t *Telegram) Update() {
	for {
		updateConfig := tgbotapi.NewUpdate(0)
		updateConfig.Timeout = 60
		updates := t.bot.GetUpdatesChan(updateConfig)
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

func fullName(user *tgbotapi.User) string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
