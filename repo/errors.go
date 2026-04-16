package repo

import "fmt"

var (
	ErrConflict = fmt.Errorf("resource already exists")
	ErrNotFound = fmt.Errorf("resource not found")
)
