package cache_test

import (
	"github.com/stretchr/testify/require"
	"github.com/trad3r/lrucache/cache"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lruCache := cache.NewCache(10)

	lruCacheWithTest := cache.NewCache(5)
	lruCacheWithTest.Add("test", "test")

	testCases := []struct {
		name    string
		cache   cache.LRUCache
		usedKey string
		method  string
		ok      bool
	}{
		{
			name:    "success add",
			cache:   lruCache,
			usedKey: "first",
			method:  "add",
			ok:      true,
		},
		{
			name:    "success get",
			cache:   lruCache,
			usedKey: "first",
			method:  "get",
			ok:      true,
		},
		{
			name:    "success remove",
			cache:   lruCache,
			usedKey: "first",
			method:  "remove",
			ok:      true,
		},
		{
			name:    "fail get",
			cache:   lruCache,
			usedKey: "first",
			method:  "get",
			ok:      false,
		},
		{
			name:    "yet isset key",
			cache:   lruCacheWithTest,
			usedKey: "test",
			method:  "add",
			ok:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			switch tc.method {
			case "get":
				_, ok := tc.cache.Get(tc.usedKey)
				require.Equal(t, tc.ok, ok)
			case "add":
				ok := tc.cache.Add(tc.usedKey, tc.usedKey)
				require.Equal(t, tc.ok, ok)
			case "remove":
				ok := tc.cache.Remove(tc.usedKey)
				require.Equal(t, tc.ok, ok)
			}
		})
	}
}
