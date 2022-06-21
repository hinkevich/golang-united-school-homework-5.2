package cache

import "testing"

func TestPut(t *testing.T) {
	cache := NewCache()
	cache.Put("one", "firs value")
	cache.Put("two", "sekond value")
	cache.Put("one", "ha-ha-ha")
	got, _ := cache.Get("one")
	if len(cache.cachemap) != 2 {
		t.Error("incorrect len")
	}
	if got != "ha-ha-ha" {
		t.Errorf("we want: %s but got: %s", "ha-ha-ha", got)
	}
}
