package domain

import "fmt"

var (
	ErrUnexpected        = fmt.Errorf("internal error")
	ErrEmptyParam = fmt.Errorf("param cannot be empty")
	ErrNotFound          = fmt.Errorf("entity not found")
	ErrEmailAlreadyExist = fmt.Errorf("entity already exist")
)
