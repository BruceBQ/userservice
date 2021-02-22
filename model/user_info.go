package model

import (
	"encoding/json"
	"io"
	"regexp"
)

type UserInfo struct {
	Name      string
	Phone     string
	Workplace string
	UpdatedAt int64
}

func UserInfoFromJSON(data io.Reader) *UserInfo {
	var u *UserInfo
	json.NewDecoder(data).Decode(&u)
	return u
}

func (u *UserInfo) Validate() map[string]interface{} {
	err := make(map[string]interface{})

	if len(u.Name) == 0 {
		err["name"] = "Nhập họ tên !"
	}

	if match, _ := regexp.MatchString(`(03|07|08|09|01[2|6|8|9])+([0-9]{8})\b`, u.Phone); !match {
		err["phone"] = "Số điện thoại không hợp lệ."
	}

	if len(u.Workplace) == 0 {
		err["workplace"] = "Nhập cơ quan !"
	}

	return err
}

func (u *UserInfo) PreUpdate() {
	u.UpdatedAt = GetMillis()
}
