package cache

import (
	"time"
)

type Notice struct {
	notice      string
	lifeTimeEnd time.Time
}

type Cache struct {
	cachemap map[string]Notice
}

func NewCache() Cache {
	var result Cache
	result.cachemap = make(map[string]Notice)
	return result
}

func (c Cache) Get(key string) (string, bool) {
	if c.cachemap == nil {
		return "", false
	}

	for keyMap, msg := range c.cachemap {
		if keyMap == key {
			if msg.lifeTimeEnd.After(time.Now()) {
				return msg.notice, true
			} else {
				delete(c.cachemap, key)
				return "", false
			}

		}

	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.cachemap[key] = Notice{notice: value, lifeTimeEnd: time.Now().Add(time.Hour)}
}

func (c Cache) Keys() []string {

	var resultarray []string
	for key := range c.cachemap {
		resultarray = append(resultarray, key)
	}
	return resultarray
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.cachemap[key] = Notice{notice: value, lifeTimeEnd: deadline}
}
