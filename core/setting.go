package core

import (
	"context"
	"fmt"
	"mofu/ent"
	"mofu/ent/setting"
	"strconv"
	"time"
)

type Setting struct {
	cache map[string]string
	db    *ent.SettingClient
}

func NewSetting(db *ent.SettingClient) *Setting {
	s := &Setting{
		db: db,
	}
	err := s.load()
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Setting) load() error {
	settings, err := s.db.Query().All(context.Background())
	if err == nil {
		s.cache = make(map[string]string, len(settings))
		for _, profile := range settings {
			s.cache[profile.Key] = profile.Value
		}

	}
	return err
}

var Defaults = map[string]string{
	SendingGap:            fmt.Sprintf("%d", 15*60),
	UpdateGap:             fmt.Sprintf("%d", 15*60),
	WaitBeforeDecided:     fmt.Sprintf("%d", 15*60),
	NumberSendingTogether: fmt.Sprintf("%d", 1),
	EnoughOrNoting:        fmt.Sprintf("%t", true),
}

const (
	SendingGap            = "sending_gap"
	UpdateGap             = "update_gap"
	WaitBeforeDecided     = "wait_before_decided"
	NumberSendingTogether = "number_sending_together"
	EnoughOrNoting        = "enough_or_nothing"
)

func (s *Setting) GetDuration(key string, dft time.Duration) time.Duration {
	second := dft
	if gap, ok := s.cache[key]; ok {
		record, err := strconv.Atoi(gap)
		if err == nil {
			second = time.Duration(record)
		}
	}
	return time.Second * second
}

func (s *Setting) GetInt(key string, dft int) int {
	if value, ok := s.cache[key]; ok {
		record, err := strconv.Atoi(value)
		if err == nil {
			return record
		}
	}
	return dft
}
func (s *Setting) GetBool(key string, dft bool) bool {
	if value, ok := s.cache[key]; ok {
		return value == "true"
	}
	return dft
}

func (s *Setting) Get(key string, dft string) string {
	if value, ok := s.cache[key]; ok {
		return value
	}
	return dft
}

func (s *Setting) Set(key string, value string) error {
	ok, err := s.db.Query().Where(setting.Key(key)).Exist(context.Background())
	if err != nil {
		return err
	}
	if !ok {
		err = s.db.Create().SetKey(key).SetValue(value).Exec(context.Background())
	} else {
		err = s.db.Update().Where(setting.Key(key)).SetValue(value).Exec(context.Background())
	}
	if err != nil {
		return err
	}
	s.cache[key] = value
	return nil
}

func (s *Setting) Remove(key string) error {
	_, err := s.db.Delete().Where(setting.Key(key)).Exec(context.Background())
	if err != nil {
		return err
	}
	delete(s.cache, key)
	return nil
}
