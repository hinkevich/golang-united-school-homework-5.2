package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {

	return Cache{}
}

func (c Cache) Get(key string) (string, bool) {
	for k, _ := range c.m {
		if k == key {
			return key, true
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
