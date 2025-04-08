package domain

import "errors"

var (
	ErrNoRecord        = errors.New("record not found")
	ErrDuplicateRecord = errors.New("duplicate record")
	ErrEmptyTitle = errors.New("title cannot empty")
)