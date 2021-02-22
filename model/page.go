package model

type Page struct {
	ID          string                 `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string                 `json:"name" bson:"name"`
	DisplayName string                 `json:"displayName" bson:"displayName"`
	Path        string                 `json:"path" bson:"path"`
	Description string                 `json:"description" bson:"description"`
	IsSettings  bool                   `json:"isSettings" bson:"isSettings"`
	Permissions map[string]interface{} `json:"permissions,omitempty" bson:"permissions,omitempty"`
}
