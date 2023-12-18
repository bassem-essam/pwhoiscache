package server

import (
	"strconv"
	"strings"
)

func StripInput(s string) string {
	return s[:len(s)-2]
}

func isValidIP(ip string) bool {
	parts := strings.Split(ip, ".")

	if len(parts) != 4 {
		return false
	}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return false
		}

		if num < 0 || num > 255 {
			return false
		}
	}

	return true
}

func isTerminated(input string) bool {
	return input[len(input)-2:] == "\r\n"
}
