package cache

import (
	"time"
)

type Cacher interface {
	Open() error
	Put(key string, value interface{}, lifetime time.Duration) error
	Del(key string) error
	Get(key string) (interface{}, error)
	Close() error
}
