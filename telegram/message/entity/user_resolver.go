package entity

import "github.com/beeper/td/tg"

// UserResolver is callback for resolving InputUser by ID.
type UserResolver = func(id int64) (tg.InputUserClass, error)
