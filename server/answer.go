package server

import (
	"fmt"

	"github.com/bassem-essam/pwhoiscache/client"
)

var INVALID_INPUT = "Error: Sorry, I don't like your input. You may ask for 'help'\n\n"
var HELP_MESSAGE = "help - this help message\r\n" +
	"<ip> - get whois of this ip\r\n" +
	"please only ask for one ip at a time\r\n" +
	"NOTE: ipv6 and bulk queries are not supported yet\r\n"

var cl *client.Client = client.New("whois.pwhois.org:43")

func GetAnswer(input string) string {
	value, ok := LookupPrefix(input)
	if ok {
		fmt.Println("Cache hit!")
		return value
	}

	fmt.Println("Cache miss!")
	result := QueryServer(input)
	StorePrefix(input, result)

	return result
}

func QueryServer(input string) string {
	// fmt.Println("Querying server...")
	return cl.Query(input)
}
