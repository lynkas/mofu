package value

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

type IMapResult interface {
	Set(key, val string)
}

type operationResultUnit struct {
	key   string
	value string
}

type operationResults struct {
	data []*operationResultUnit
	*text
}

func NewOperationResults(title string) *operationResults {
	return &operationResults{
		text: NewText(title),
	}
}

func (a *operationResults) Set(key, val string) {
	for _, datum := range a.data {
		if datum.key == key {
			datum.value = val
			return
		}
	}
	a.data = append(a.data, &operationResultUnit{
		key:   key,
		value: val,
	})
}

func (a *operationResults) Content() string {
	result := strings.Builder{}
	result.WriteString(fmt.Sprintf("%s\n", a.text.Content()))
	maxLen := 0
	for _, unit := range a.data {
		if len(unit.key) > maxLen {
			maxLen = len(unit.key)
		}
	}

	for _, unit := range a.data {
		spaceCount := maxLen - len(unit.key) + 1
		result.WriteString(fmt.Sprintf("<code>%s</code><code>%s</code><code>%s</code>\n", unit.key, strings.Repeat(" ", spaceCount), unit.value))
	}
	return result.String()
}
func (a *operationResults) ImageURL() string {
	return ""
}

func (a *operationResults) Message(channelID int64, replyID int) tgbotapi.Chattable {
	return TextMessage(channelID, replyID, a)
}

func (a *operationResults) EditMessage(chatID int64, replyID int) tgbotapi.Chattable {
	return EditMessage(chatID, replyID, a)
}
