package cache

import (
	"container/list"
	"sync"
	"time"
)

type LRU struct {
	lock          sync.RWMutex
	size          int
	len           int
	evictList     *list.List
	items         map[string]*list.Element
	defaultExpiry time.Duration
	name          string
}

type LRUOptions struct {
	Name          string
	Size          int
	DefaultExpiry time.Duration
}

// entry is used to hold a value in the evictList.
type entry struct {
	key     string
	value   interface{}
	expires time.Time
}

func NewLRU(opts LRUOptions) Cache {
	return &LRU{
		name:          opts.Name,
		size:          opts.Size,
		evictList:     list.New(),
		items:         make(map[string]*list.Element, opts.Size),
		defaultExpiry: opts.DefaultExpiry,
	}
}
func (l *LRU) Purge() error {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.len = 0
	return nil
}

func (l *LRU) Set(key string, value interface{}) error {
	return l.SetWithExpiry(key, value, 0)
}

func (l *LRU) SetWithDefaultExpiry(key string, value interface{}) error {
	return l.SetWithExpiry(key, value, l.defaultExpiry)
}

func (l *LRU) SetWithExpiry(key string, value interface{}, ttl time.Duration) error {
	return l.set(key, value, ttl)
}

func (l *LRU) Get(key string, value interface{}) error {
	return l.get(key, value)
}

func (l *LRU) Remove(key string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	if ent, ok := l.items[key]; ok {
		l.removeElement(ent)
	}
	return nil
}

func (l *LRU) Keys() ([]string, error) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	keys := make([]string, l.len)
	i := 0
	for ent := l.evictList.Back(); ent != nil; ent = ent.Prev() {
		e := ent.Value.(*entry)
		keys[i] = e.key
		i++
		// if e.generation == l.currentGeneration {
		// 	keys[i] = e.key
		// 	i++
		// }
	}
	return keys, nil
}

func (l *LRU) Len() (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	return l.len, nil
}

// Name returns the name of the cache
func (l *LRU) Name() string {
	return l.name
}

func (l *LRU) set(key string, value interface{}, ttl time.Duration) error {
	var expires time.Time

	if ttl > 0 {
		expires = time.Now().Add(ttl)
	}

	// var buf []byte
	// var err error

	// if msgpVal, ok := value.(msgp.Marshaler); ok {
	// 	buf, err = msgpVal.MarshalMsg(nil)
	// }
	l.lock.Lock()
	defer l.lock.Unlock()

	// Check for existing item, ignoring expiry since we'd update anyway.
	if ent, ok := l.items[key]; ok {
		l.evictList.MoveToFront(ent)
		e := ent.Value.(*entry)
		e.value = value
		e.expires = expires
		l.len++
		return nil
	}

	// Add new item
	ent := &entry{key, value, expires}
	entry := l.evictList.PushFront(ent)
	l.items[key] = entry
	l.len++

	if l.evictList.Len() > l.size {
		l.removeElement(l.evictList.Back())
	}
	return nil
}

func (l *LRU) get(key string, value interface{}) error {
	_, err := l.getItem(key)
	if err != nil {
		return err
	}

	return nil
}

func (l *LRU) getItem(key string) (interface{}, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	ent, ok := l.items[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	e := ent.Value.(*entry)
	if !e.expires.IsZero() && time.Now().After(e.expires) {
		l.removeElement(ent)
		return nil, ErrKeyNotFound
	}
	l.evictList.MoveToFront(ent)
	return e.value, nil
}

func (l *LRU) removeElement(e *list.Element) {
	l.evictList.Remove(e)
	kv := e.Value.(*entry)
	l.len--
	delete(l.items, kv.key)
}
