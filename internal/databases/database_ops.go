package databases

import "fmt"

type RoleOps interface {
	create() string
	update() string
	delete() string
}

type UserOps interface {
	create() string
	update() string
	delete() string
}

type A struct {
}

func (a A) create() string {
	fmt.Printf("%s is working\n", a)
	return ""
}
