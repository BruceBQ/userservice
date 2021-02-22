package model

import (
	"encoding/json"
	"math/rand"
	"time"
)

const (
	SESSION_CACHE_SIZE       = 10000
	SESSION_ACTIVITY_TIMEOUT = 1000 * 60 * 5
)

type Session struct {
	ID             string                 `json:"id,omitempty" bson:"_id,omitempty"`
	Token          string                 `json:"token" bson:"token"`
	CreatedAt      int64                  `json:"createdAt" bson:"created_at"`
	ExpiresAt      int64                  `json:"expiresAt" bson:"expires_at"`
	LastActivityAt int64                  `json:"lastActivityAt" bson:"last_activity_at"`
	DeviceID       string                 `json:"deviceId" bson:"device_id"`
	UserID         string                 `json:"userId" bson:"user_id"`
	IsAuth         bool                   `json:"isAuth" bson:"is_auth"`
	Props          map[string]interface{} `json:"props" bson:"props"`
}

func (s *Session) ToJSON() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s *Session) PreSave() {

	if s.Token == "" {
		s.Token = NewToken(40)
	}
	if s.CreatedAt == 0 {
		s.CreatedAt = GetMillis()
	}
	s.LastActivityAt = s.CreatedAt

	if s.Props == nil {
		s.Props = make(map[string]interface{})
	}
}

func (s *Session) IsExpired() bool {
	if s.ExpiresAt <= 0 {
		return false
	}

	if GetMillis() > s.ExpiresAt {
		return true
	}

	return false
}

func NewToken(size int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, size)
	for i := range b {
		b[i] = ALL_CHARACTER[seededRand.Intn(len(ALL_CHARACTER))]
	}
	return string(b)
}
