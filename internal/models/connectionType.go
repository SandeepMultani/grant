package models

import "fmt"

type ConnectionType int

const (
	Postgres ConnectionType = iota + 1
	MySql
)

func (c ConnectionType) String() string {
	connectionTypes := [...]string{"postgres", "mysql"}
	if c < Postgres || c > MySql {
		return fmt.Sprintf("Invalid ConnectionType(%d)", int(c))
	}
	return connectionTypes[c-1]
}

func (c ConnectionType) IsValid() bool {
	switch c {
	case Postgres, MySql:
		return true
	}
	return false
}
