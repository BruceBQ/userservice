package model

import (
	"encoding/json"
	"io"
)

const (
	SUPER_ADMIN_ROLE = "Super Admin"

	ROLE_NAME_MAX_LENGTH         = 128
	ROLE_DISPLAY_NAME_MAX_LENGTH = 128
)

type Role struct {
	ID          string                 `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string                 `json:"name" bson:"name"`
	Description string                 `json:"description" bson:"description"`
	Pages       map[string]interface{} `json:"pages,omitempty" bson:"pages,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty" bson:"permissions,omitempty"`
	BuiltIn     bool                   `json:"builtin,omitempty" bson:"builtin,omitempty"`
	CreatedAt   int64                  `json:"createdAt,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   int64                  `json:"updatedAt,omitempty" bson:"updated_at,omitempty"`
	DeletedAt   int64                  `json:"deletedAt,omitempty" bson:"deleted_at,omitempty"`
}

type RoleSlice []*Role

func RoleFromJSON(data io.Reader) *Role {
	var r *Role
	json.NewDecoder(data).Decode(&r)
	return r
}

func (r *Role) Validate() map[string]interface{} {
	errs := make(map[string]interface{})

	if len(r.Name) == 0 {
		errs["name"] = "Tên nhóm người dùng không được để trống."
	}

	if len(r.Description) == 0 {
		errs["description"] = "Mô tả không được để trống."
	}

	if len(r.Permissions) == 0 {
		errs["permissions"] = "Chọn API."
	}

	return errs
}

func (r *Role) PreSave() {
	if r.CreatedAt == 0 {
		r.CreatedAt = GetMillis()
	}

	r.UpdatedAt = r.CreatedAt
	r.BuiltIn = false
}

func (r *Role) PreUpdate() {
	r.UpdatedAt = GetMillis()
}

func (r *Role) ToJSON() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (rs *RoleSlice) ToJSON() string {
	b, _ := json.Marshal(rs)
	return string(b)
}
