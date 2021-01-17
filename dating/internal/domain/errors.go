package domain

import "fmt"

var (
	ErrUnexpected        = fmt.Errorf("internal error")
	ErrNotFound          = fmt.Errorf("entity not found")
	ErrEmailAlreadyExist = fmt.Errorf("entity already exist")
)
