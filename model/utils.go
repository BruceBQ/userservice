package model

import (
	"encoding/json"
	"io"
	"net/mail"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	LOWERCASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
	ALL_CHARACTER     = LOWERCASE_LETTERS + UPPERCASE_LETTERS + NUMBERS
	SYMBOLS           = " !\"\\#$%&'()*+,-./:;<=>?@[]^_`|~"
)

type StringInterface map[string]interface{}
type StringMap map[string]string

type AppError struct {
	Id            string                 `json:"id,omitempty"`
	Message       string                 `json:"message"`
	DetailedError string                 `json:"detailedError,omitempty"`
	Where         string                 `json:"path,omitempty"`
	StatusCode    int                    `json:"statusCode,omitempty"`
	Params        map[string]interface{} `json:"params,omitempty"`
}

// MapToJson converts a map to a json string
func MapToJson(objmap map[string]string) string {
	b, _ := json.Marshal(objmap)
	return string(b)
}

// MapFromJson will decode the key/value pair map
func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}

func NewAppError(where string, id string, message string, params map[string]interface{}, details string, status int) *AppError {
	ap := &AppError{}
	ap.Id = id
	ap.Message = message
	ap.DetailedError = details
	ap.Where = where
	ap.StatusCode = status
	ap.Params = params
	return ap
}

func (er *AppError) ToJson() string {
	b, _ := json.Marshal(er)
	return string(b)
}

func (er *AppError) Sanitize(options map[string]bool) {
	// er.Id = ""
	// er.Where = ""
	// er.StatusCode = 0
	// er.DetailedError = ""
}

func StringInterfaceToJson(objmap map[string]interface{}) string {
	b, _ := json.Marshal(objmap)
	return string(b)
}

func StringToJson(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func IsValidString(s string) bool {
	if len(s) == 0 {
		return false
	}

	return true
}

func IsLower(s string) bool {
	return strings.ToLower(s) == s
}

func IsValidEmail(email string) bool {
	if !IsLower(email) {
		return false
	}

	if addr, err := mail.ParseAddress(email); err != nil {
		return false
	} else if addr.Name != "" {
		return false
	}

	return true
}

func IsMongoId(id string) bool {
	_, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return false
	}

	return true
}

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
