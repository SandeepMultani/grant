package models

type Role struct {
	Name      string   `json:"name"`
	Type      RoleType `json:"type"`
	Grants    []string `json:"grants"`
	Databases []string `json:"databases"`
}
