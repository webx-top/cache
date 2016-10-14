package cache

import (
	"errors"
)

var (
	ErrExpired = errors.New(`The data has expired`)
)
