package model

import "encoding/json"

type AuditSearch struct {
	AuditLogEntries []*Audit      `json:"auditLogEntries"`
	Users           []*User       `json:"users,omitempty"`
	Permissions     []*Permission `json:"permissions,omitempty"`
	Page            *int          `json:"page"`
	PerPage         int           `json:"perPage"`
	TotalPage       int64         `json:"totalPage"`
}

func (a *AuditSearch) ToJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

type AuditFilter struct {
	Users       []*User       `json:"users"`
	Permissions []*Permission `json:"permissions"`
}

func (f *AuditFilter) ToJson() string {
	b, _ := json.Marshal(f)
	return string(b)
}
