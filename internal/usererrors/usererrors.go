package usererrors

import "errors"

var (
	ErrIncorrectNumberOfArgs = errors.New("incorrect number of arguments")
	ErrStepsNumber           = errors.New("number of steps less than zero")
	ErrHeight                = errors.New("incorrect height")
	ErrWeight                = errors.New("incorrect weight")
	ErrDuration              = errors.New("incorrect duration")
	ErrActivity              = errors.New("unknown activity")
)
