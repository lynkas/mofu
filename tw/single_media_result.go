package tw

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type SingleMediaResult struct {
	Original  *AuthorTweet `json:"original"`
	BroadCast *Author      `json:"broadcast"`
	Media     *Media       `json:"media"`
}

func SingleMediaResultFrom(data []byte) *SingleMediaResult {
	var result SingleMediaResult
	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Warn("SingleMediaResultFrom ", err)
	}
	return &result
}
func (s SingleMediaResult) Dump() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		log.Warn(err)
	}
	return data
}

func (s *SingleMediaResult) Key() string {
	return s.Media.Key
}

func (s *SingleMediaResult) Type() string {
	return s.Media.Type
}

func (s *SingleMediaResult) URL() string {
	return s.Media.URL
}

func (s *SingleMediaResult) Tweet() ITweet {
	return s.Original.Tweet
}

func (s *SingleMediaResult) Author() IAuthor {
	return s.Original.Author
}

func (s *SingleMediaResult) Broadcast() IAuthor {
	return s.BroadCast
}
