package tg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	olog "log"
	"mofu/utils"
	"strings"
	"time"
)

type heartbeat struct {
	*utils.PeriodFrequency
	messageID int
}

func NewHeartbeat() *heartbeat {
	return &heartbeat{
		PeriodFrequency: utils.NewFrequencyLimit(1, time.Minute*1),
	}
}

func (t *Telegram) heartbeat() {
	h := NewHeartbeat()
	for {
		msg, err := t.Log(fmt.Sprintf("Heartbeat init"), false)
		if err != nil {
			olog.Fatal(err)
		}

	inner:
		for {
			h.Can()

			select {
			case <-t.heartbeatBreak:
				break inner
			default:
				err := t.EditLog(fmt.Sprintf("Heartbeat:\n%s", time.Now().Format(time.RFC3339)), msg.MessageID)
				if err != nil {
					olog.Println(err)
					break inner
				}
			}
		}
	}
}

func (t *Telegram) EditLog(content string, msgID int) error {
	msg := tgbotapi.NewEditMessageText(t.warningRoomID, msgID, content)
	_, err := t.Request(msg)
	if err != nil {
		olog.Println(err)
	}
	return err
}
func (t *Telegram) Log(content string, pin bool) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(t.warningRoomID, content)
	msg.DisableNotification = true
	warn := strings.Contains(strings.ToLower(content), "error") || strings.Contains(strings.ToLower(content), "warn")
	msg.DisableNotification = !warn
	if warn {
		select {
		case t.heartbeatBreak <- 1:
		default:
		}
	}
	m, err := t.Send(msg)
	if err != nil {
		olog.Println(err)
		return m, err
	}
	if pin {
		_, err = t.bot.Request(tgbotapi.PinChatMessageConfig{
			ChatID:              t.warningRoomID,
			MessageID:           m.MessageID,
			DisableNotification: false,
		})
		if err != nil {
			olog.Println(err)
		}

	}
	return m, err
}

func (t *Telegram) Write(p []byte) (n int, err error) {
	return t.logBuffer.Write(p)
}

func (t *Telegram) UpdateLog() {
	go t.heartbeat()
	cache := strings.Builder{}
	retry := 0
	for {
		b := t.logBuffer.Next(1)
		if len(b) == 0 {
			<-time.After(time.Second)
			continue
		}
		if b[0] == 0 {
			continue
		}
		if b[0] == '\n' {
			_, err := t.Log(cache.String(), true)
			if err != nil {
				retry += 1
				<-time.After(time.Second)
				if retry >= 5 {
					olog.Fatal(err)
				}
				continue
			} else {
				cache.Reset()
			}
		} else {
			cache.Write(b)
		}
	}
}
