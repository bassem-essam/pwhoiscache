package server

import "sync"

// this is legacy and test-only cache

var cache sync.Map = sync.Map{}

func Lookup(input string) (string, bool) {
	value, ok := cache.Load(input)
	if ok {
		return value.(string), true
	}

	return "", false
}

func Store(input, value string) {
	cache.Store(input, value)
}
