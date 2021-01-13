package domain

import "fmt"

var Error struct {
	data map[string]interface{}
}

var (
	ErrUnexpected = fmt.Errorf("internal error")
	ErrNotFound   = fmt.Errorf("entity not found")
)
