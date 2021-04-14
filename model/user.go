package model

import (
	"encoding/json"
	"io"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	ME = "@me"

	USER_PASSWORD_MIN_LENGTH = 8
	USER_PASSWORD_MAX_LENGTH = 72
)

type User struct {
	ID                 string                 `json:"id,omitempty" bson:"_id,omitempty"`
	Email              string                 `json:"email,omitempty"`
	Phone              string                 `json:"phone,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Description        string                 `json:"description,omitempty"`
	Password           string                 `json:"password,omitempty" bson:"password,omitempty"`
	Workplace          string                 `json:"workplace,omitempty" bson:"workplace,omitempty"`
	RoleID             string                 `json:"roleId,omitempty" bson:"role_id,omitempty"`
	Role               *Role                  `json:"role,omitempty" bson:"role,omitempty"`
	Cameras            []string               `json:"cameras,omitempty" bson:"cameras,omitempty"`
	CameraDetail       map[string]interface{} `json:"cameraDetail,omitempty" bson:"cameraDetail,omitempty"`
	LastPasswordUpdate int64                  `json:"lastPasswordUpdate,omitempty" bson:"last_password_update,omitempty"`
	CreatedAt          int64                  `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt          int64                  `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
	Token              string                 `json:"token,omitempty"`
	BuiltIn            bool                   `json:"builtin,omitempty" bson:"builtin,omitempty"`
}

type UserSlice []*User

func UserFromJSON(data io.Reader) *User {
	var u *User
	json.NewDecoder(data).Decode(&u)
	return u
}

type UserWithSession struct {
	User  *User  `json:"user,omitempty"`
	Token string `json:"token"`
}

type UserWithPage struct {
	User *User   `json:"user"`
	Page []*Page `json:"page"`
}

func (u *User) Validate() map[string]interface{} {
	err := make(map[string]interface{})

	if !IsValidEmail(u.Email) {
		err["email"] = "Email không hợp lệ !"
	}

	if len(u.Email) == 0 {
		err["email"] = "Nhập email !"
	}

	if len(u.Name) == 0 {
		err["name"] = "Nhập họ tên !"
	}

	if len(u.Workplace) == 0 {
		err["workplace"] = "Nhập cơ quan !"
	}

	if match, _ := regexp.MatchString(`(03|07|08|09|01[2|6|8|9])+([0-9]{8})\b`, u.Phone); !match {
		err["phone"] = "Số điện thoại không hợp lệ."
	}

	if !IsMongoId(u.RoleID) {
		err["role"] = "Nhóm người dùng không hợp lệ !"
	}

	if len(u.Cameras) == 1 {
		if u.Cameras[0] != "*" && !IsMongoId(u.Cameras[0]) {
			err["cameras"] = "Danh sách camera không hợp lệ !"
		}
	} else {
		for _, v := range u.Cameras {
			if !IsMongoId(v) {
				err["cameras"] = "Danh sách camera không hợp lệ !"
				break
			}
		}
	}

	return err
}

// ToJson convert a User to a json string
func (u *User) ToJson() string {
	b, _ := json.Marshal(u)
	return string(b)
}

// Remove any private data from the user object
func (u *User) Sanitize(options map[string]bool) {
	u.Password = ""

	if len(options) != 0 && !options["cameras"] {
		u.Cameras = []string{}
	}
}

func (u *User) PreSave() {
	u.CreatedAt = GetMillis()
	u.UpdatedAt = u.CreatedAt
	u.LastPasswordUpdate = u.CreatedAt
	u.Password = HashPassword(u.Email)
}

func (u *User) PreUpdate() {
	u.UpdatedAt = GetMillis()
}

func (us *UserWithSession) ToJSON() string {
	b, _ := json.Marshal(us)
	return string(b)
}

func (up *UserWithPage) ToJSON() string {
	b, _ := json.Marshal(up)
	return string(b)
}

// HashPassword generates a hash using the bcrypt.GenerateFromPassword
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

// ComparePassword compares the hash
func ComparePassword(hash string, password string) bool {
	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
