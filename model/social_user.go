package model

import (
	"encoding/json"
	"io"
)

const (
	GOOGLE   = "google"
	FACEBOOK = "facebook"
	APPLE    = "apple"
)

type SocialUser struct {
	Id           string    `json:"id" bson:"_id,omitempty"`
	AuthType     string    `json:"authType" bson:"authType,omitempty"`
	DisplayName  string    `json:"displayName" bson:"displayName,omitempty"`
	Email        string    `json:"email" bson:"email"`
	Username     string    `json:"username" bson:"username"`
	Password     string    `json:"password,omitempty"`
	PhoneNumber  string    `json:"phoneNumber" bson:"phoneNumber"`
	Avatar       string    `json:"avatar"`
	PlateNumbers []*string `json:"plateNumbers" bson:"plateNumbers"`
	CreatedAt    int64     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt    int64     `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

type SocialUserWithSession struct {
	User  *SocialUser `json:"user"`
	Token string      `json:"token"`
}

func SocialUserFromJson(data io.Reader) *SocialUser {
	var u *SocialUser
	json.NewDecoder(data).Decode(&u)
	return u
}

func (u *SocialUser) Validate() map[string]interface{} {
	errs := make(map[string]interface{})

	if u.AuthType != GOOGLE && u.AuthType != FACEBOOK && u.AuthType != APPLE {
		errs["authType"] = "Loại xác thực không hợp lệ"
	}

	if !IsValidEmail(u.Email) {
		errs["email"] = "Email không hợp lệ"
	}

	return errs
}

func (u *SocialUser) PreSave() {
	if u.CreatedAt == 0 {
		u.CreatedAt = GetMillis()
		u.UpdatedAt = u.CreatedAt
	}
}

func (u *SocialUser) ToJson() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *SocialUser) Sanitize() {
	u.PlateNumbers = nil
}

func (u *SocialUserWithSession) ToJson() string {
	b, _ := json.Marshal(u)
	return string(b)
}
