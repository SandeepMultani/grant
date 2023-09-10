package models

type Connection struct {
	Type             ConnectionType `json:"type"`
	ConnectionString string         `json:"connectionString"`
}
