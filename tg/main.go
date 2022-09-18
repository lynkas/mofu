package tg

import (
	"bytes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"mofu/utils"
	"time"
)

type Telegram struct {
	bot                  *tgbotapi.BotAPI
	controlRoomID        int64
	warningRoomID        int64
	broadcastRoomID      int64
	frequencyLimit       *utils.PeriodFrequency
	logBuffer            *bytes.Buffer
	admin                []tgbotapi.ChatMember
	adminCacheExpireTime time.Time
	heartbeatBreak       chan int
	function             ITelegramFunc
}

func New(token string, ControlRoomID, BroadcastRoomID, WarningRoomID int64, coreFunc ITelegramFunc) *Telegram {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	limit := utils.NewFrequencyLimit(10, time.Second*10)
	return &Telegram{
		bot:             bot,
		controlRoomID:   ControlRoomID,
		broadcastRoomID: BroadcastRoomID,
		warningRoomID:   WarningRoomID,
		frequencyLimit:  limit,
		logBuffer:       bytes.NewBuffer(make([]byte, 9999)),
		heartbeatBreak:  make(chan int, 1),
		function:        coreFunc,
	}
}

func (t *Telegram) admins() []tgbotapi.ChatMember {
	if t.adminCacheExpireTime.After(time.Now()) {
		return t.admin
	}

	admin, err := t.bot.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		ChatConfig: tgbotapi.ChatConfig{ChatID: t.controlRoomID},
	})
	if err != nil {
		log.Warn(err)
		return nil
	}

	t.adminCacheExpireTime = time.Now().Add(time.Minute * 5)
	t.admin = admin
	return t.admin
}

func (t *Telegram) isAdmin(sender *tgbotapi.User) bool {
	for _, member := range t.admins() {
		if member.User.ID == sender.ID {
			return true
		}
	}
	log.Warn("this guy is not an admin, ", sender)
	return false
}

func (t *Telegram) validCommand(update tgbotapi.Update) bool {
	return t.controlRoomID == update.Message.Chat.ID && (update.Message != nil && t.isAdmin(update.Message.From) || update.CallbackQuery != nil && t.isAdmin(update.CallbackQuery.From))
}

func (t *Telegram) Send(msg tgbotapi.Chattable) (tgbotapi.Message, error) {
	t.frequencyLimit.Can()
	return t.bot.Send(msg)
}

func (t *Telegram) Request(msg tgbotapi.Chattable) (*tgbotapi.APIResponse, error) {
	t.frequencyLimit.Can()
	return t.bot.Request(msg)
}
