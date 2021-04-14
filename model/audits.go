package model

import (
	"encoding/json"
	"io"
)

type Audits []*Audit

func (o Audits) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func AuditsFromJSON(data io.Reader) Audits {
	var o Audits
	json.NewDecoder(data).Decode(&o)
	return o
}
