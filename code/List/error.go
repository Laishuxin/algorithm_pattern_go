package List

import "errors"

var (
	ErrEmptyList = errors.New("list: list is empty!")
	ErrIndexOutOfRange = errors.New("list: index out of range!")
	ErrHasNotNext = errors.New("list iterator: iterator has not next!")
)
