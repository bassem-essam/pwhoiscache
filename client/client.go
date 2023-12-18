package client

import (
	"io"
	"net"
)

type Client struct {
	WhoisServer string
}

func New(whoisServer string) *Client {
	return &Client{whoisServer}
}

func (c *Client) Query(q string) string {
	client, err := net.Dial("tcp", c.WhoisServer)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	// fmt.Printf("Sending query... (%v) to %s\n", []byte(q), c.Server)
	client.Write([]byte(q + "\r\n"))
	result, err := io.ReadAll(client)
	if err != nil {
		panic(err)
	}

	return string(result)
}
