package models

import "fmt"

type RoleType int

const (
	Database RoleType = iota + 1
	Schema
	Table
)

func (c RoleType) String() string {
	roleTypes := [...]string{"postgres", "mysql"}
	if c < Database || c > Table {
		return fmt.Sprintf("Invalid RoleType(%d)", int(c))
	}
	return roleTypes[c-1]
}

func (c RoleType) IsValid() bool {
	switch c {
	case Database, Schema, Table:
		return true
	}
	return false
}
