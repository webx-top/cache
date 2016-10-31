package ttlmap

import (
	"runtime"
	"time"

	"github.com/admpub/go-ttlmap"
	"github.com/webx-top/cache"
)

func New(cap int, hooks ...func(string, *ttlmap.Item)) cache.Cacher {
	r := &TTLMap{
		Options: &ttlmap.Options{
			InitialCapacity: cap,
			OnWillExpire:    nil,
			OnWillEvict:     nil,
		},
	}
	switch len(hooks) {
	case 2:
		r.Options.OnWillEvict = hooks[1]
		fallthrough
	case 1:
		r.Options.OnWillExpire = hooks[0]
	}
	return r
}

type TTLMap struct {
	Options *ttlmap.Options
	m       *ttlmap.Map
}

func (t *TTLMap) Init() error {
	t.m = ttlmap.New(t.Options)
	runtime.SetFinalizer(t, func(t *TTLMap) error {
		t.m.Drain()
		return nil
	})
	return nil
}

func (t *TTLMap) Open() error {
	return nil
}

func (t *TTLMap) Put(key string, value interface{}, lifetime time.Duration) error {
	return t.m.Set(key, ttlmap.NewItemWithTTL(value, lifetime))
}

func (t *TTLMap) Del(key string) error {
	t.m.Delete(key)
	return nil
}

func (t *TTLMap) Get(key string) (interface{}, error) {
	item := t.m.Get(key)
	if item.Expired() {
		return nil, cache.ErrExpired
	}
	return item.Value, nil
}

func (t *TTLMap) Close() error {
	return nil
}
