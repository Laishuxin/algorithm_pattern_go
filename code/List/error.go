package List

import "errors"

var (
	ErrEmptyList = errors.New("list: list is empty!")
	ErrIndexOutOfRange = errors.New("list: index out of range!")
)
