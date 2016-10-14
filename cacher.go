package cache

import (
	"time"
)

type Cacher interface {
	Init() error
	Open() error
	Put(key string, value interface{}, lifetime time.Duration) error
	Del(key string) error
	Get(key string) (interface{}, error)
	Close() error
}
