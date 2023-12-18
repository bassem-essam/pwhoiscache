package main

import (
	"flag"
	"fmt"

	"github.com/bassem-essam/pwhoiscache/client"
	"github.com/bassem-essam/pwhoiscache/server"
)

var (
	whoisServer *string = flag.String("s", "whois.pwhois.org:43", "Server to connect to")
	clientMode  *bool   = flag.Bool("q", false, "use client mode")
)

func main() {
	flag.Parse()

	if *clientMode {
		c := client.New(*whoisServer)
		s := c.Query(flag.Arg(0))
		fmt.Print(s)
	} else {
		server.StartServer()
	}
}
