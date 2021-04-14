package model

import "encoding/json"

type Permission struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	DisplayName string `json:"displayName" bson:"displayName"`
	Description string `json:"description,omitempty"  bson:"description,omitempty"`
}

func (pu *Permission) ToJSON() string {
	b, _ := json.Marshal(pu)
	return string(b)
}

func (pu *Permission) Sanitize() {
	pu.ID = ""
	pu.Description = ""
}

var PERMISSION_GET_PERMISSIONS *Permission
var PERMISSION_GET_ROLES *Permission
var PERMISSION_CREATE_ROLE *Permission
var PERMISSION_UPDATE_ROLE *Permission
var PERMISSION_GET_ROLE *Permission
var PERMISSION_DELETE_ROLE *Permission

var PERMISSION_GET_USERS *Permission
var PERMISSION_CREATE_USER *Permission
var PERMISSION_GET_USER *Permission
var PERMISSION_DELETE_USER *Permission
var PERMISSION_UPDATE_USER *Permission

var PERMISSION_GET_AUDIT_LOGS *Permission

func initializePermissions() {
	PERMISSION_GET_PERMISSIONS = &Permission{
		Name: "get_permissions",
	}
	PERMISSION_GET_ROLES = &Permission{
		Name: "get_roles",
	}
	PERMISSION_CREATE_ROLE = &Permission{
		Name: "create_role",
	}
	PERMISSION_UPDATE_ROLE = &Permission{
		Name: "update_role",
	}
	PERMISSION_GET_ROLE = &Permission{
		Name: "get_single_role",
	}
	PERMISSION_DELETE_ROLE = &Permission{
		Name: "delete_role",
	}
	PERMISSION_GET_USERS = &Permission{
		Name: "get_users",
	}
	PERMISSION_CREATE_USER = &Permission{
		Name: "create_user",
	}
	PERMISSION_GET_USER = &Permission{
		Name: "get_single_user",
	}
	PERMISSION_DELETE_USER = &Permission{
		Name: "delete_user",
	}
	PERMISSION_UPDATE_USER = &Permission{
		Name: "update_user",
	}
	PERMISSION_GET_AUDIT_LOGS = &Permission{
		Name: "get_audit_logs",
	}
}

func init() {
	initializePermissions()
}
