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
	result.cacheNotices = make([]Notice, 0)
	return result
}

func (c Cache) Get(key string) (string, bool) {
	if c.cacheNotices == nil {
		return "", false
	}
	for _, notice := range c.cacheNotices {
		if notice.key == key && notice.lifeTime.After(time.Now()) {
			return notice.notice, true
		}

	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	checkAdd := false
	for _, notice := range c.cacheNotices {
		if notice.key == key {
			notice.notice = value
			notice.lifeTime.Add(time.Hour)
			checkAdd = true
			break
		}
	}
	if !checkAdd {
		c.cacheNotices = append(c.cacheNotices, Notice{key, value, time.Now().Add(time.Hour)})
	}
}

func (c Cache) Keys() []string {

	var resultarray []string
	for _, notice := range c.cacheNotices {
		resultarray = append(resultarray, notice.key)
	}
	return resultarray
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {

}
