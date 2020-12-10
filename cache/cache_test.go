package cache_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/trad3r/lrucache/cache"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lruCache := createCache(3000)

	testCases := []struct {
		name    string
		cache   cache.LRUCache
		usedKey string
		method  string
		ok      bool
	}{
		{
			name:    "success get",
			cache:   lruCache,
			usedKey: "0",
			method:  "get",
			ok:      true,
		},
		{
			name:    "success add",
			cache:   lruCache,
			usedKey: "-1",
			method:  "add",
			ok:      true,
		},
		{
			name:    "success remove",
			cache:   lruCache,
			usedKey: "-1",
			method:  "remove",
			ok:      true,
		},
		{
			name:    "fail get",
			cache:   lruCache,
			usedKey: "-1",
			method:  "get",
			ok:      false,
		},
		{
			name:    "yet isset key",
			cache:   lruCache,
			usedKey: "0",
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

func createCache(count int) *cache.Cache {
	lruCache := cache.NewCache(count)
	for i := 0; i < lruCache.Count; i++ {
		el := fmt.Sprintf("%d", i)
		lruCache.Add(el, el)
	}

	return lruCache
}
