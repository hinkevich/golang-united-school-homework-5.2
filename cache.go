package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {
	var result Cache
	result.m = make(map[string]string)

	return result
}

func (c Cache) Get(key string) (string, bool) {
	if c.m == nil {
		return "", false
	}
	for k, v := range c.m {
		if k == key {
			return v, true
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

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	if deadline.After(time.Now()) {
		c.Put(key, value)
	}
}
