package model

import (
	"encoding/json"
	"io"
	"time"
)

const (
	AuditPerPage = 20
)

type Audit struct {
	ID             string      `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         string      `json:"userId,omitempty" bson:"user_id,omitempty"`
	User           *User       `json:"user,omitempty" bson:"user,omitempty"`
	PermissionName string      `json:"permission_name,omitempty" bson:"permission_name,omitempty"`
	Permission     *Permission `json:"permission,omitempty" bson:"permission,omitempty"`
	Data           string      `json:"data" bson:"data"`
	CreatedAt      int64       `json:"createdAt" bson:"created_at"`
}

func (o *Audit) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func (o *Audit) PreSave() {
	if o.CreatedAt == 0 {
		o.CreatedAt = GetMillis()
	}

}

func AuditFromJSON(data io.Reader) *Audit {
	var o *Audit
	json.NewDecoder(data).Decode(&o)
	return o
}

type AuditRequest struct {
	Page       int        `json:"page"`
	Before     *time.Time `json:"before"`
	UserId     string     `json:"userId"`
	Permission string     `json:"permission"`
}

func AuditRequestFromJson(data io.Reader) *AuditRequest {
	var o *AuditRequest
	json.NewDecoder(data).Decode(&o)
	return o
}
