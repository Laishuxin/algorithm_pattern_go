package Queue

import "errors"

var (
	ErrEmptyQueue = errors.New("queue: empty queue")
	ErrFullQueue = errors.New("queue: full queue")
)
