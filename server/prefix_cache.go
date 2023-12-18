package server

import (
	"net"
	"strconv"
	"strings"
	"sync"
)

var prefixCache sync.Map = sync.Map{}

func LookupPrefix(input string) (string, bool) {
	// We have two options here (that I can think of)):
	// 1. We can get all mutations of the input and check if any of them is in the cache
	// 2. We can iterate over the cache and check if any of the keys is a prefix of the input

	// Option 1
	for _, prefix := range AllPrefixes(input) {
		value, ok := prefixCache.Load(prefix)
		if ok {
			return strings.Replace(value.(string), "REPLACEME", input, 1), true
		}
	}

	return "", false
}

func AllPrefixes(input string) []string {
	var prefixes []string
	for i := 32; i > 0; i-- {
		prefix := Prefix(input + "/" + strconv.Itoa(i))
		prefixes = append(prefixes, prefix)
	}

	return prefixes
}

func Prefix(ip string) string {
	_, ipnet, err := net.ParseCIDR(ip)
	if err != nil {
		panic(err)
	}

	return ipnet.String()
}

func StorePrefix(input, value string) {
	if value == "" {
		prefixCache.Store(input+"/24", "")
		return
	}
	prefix := ExtractPrefix(value)
	parts := strings.SplitN(value, "\n", 2) // IP and the rest
	parts[0] = "IP: REPLACEME"
	prefixCache.Store(prefix, strings.Join(parts, "\n"))
}

func ExtractPrefix(raw string) string {
	return strings.Split(strings.Split(raw, "\n")[2], ": ")[1]
}
