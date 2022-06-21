package cache

import "time"

type Notice struct {
	key      string
	notice   string
	lifeTime time.Time
}

type Cache struct {
	cacheNotices []Notice
}

func NewCache() Cache {
	var result Cache
	//result.cacheNotices := make([]Notice)

	return result
}

func (c Cache) Get(key string) (string, bool) {
	if c.cacheNotices == nil {
		return "", false
	}
	for _, notice := range c.cacheNotices {
		if notice.key == key {
			return notice.notice, true
		}

	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.m[key] = value
}

func (c Cache) Keys() []string {

	var resultarray []string
	for key, _ := range c.m {
		resultarray = append(resultarray, key)
	}
	return resultarray
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
}
