package server

import (
	"bufio"
	"fmt"
	"net"
)

func StartServer() {
	server, err := net.Listen("tcp", ":43")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on port 43...")

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	sc := bufio.NewReader(conn)
	var answer string

	input, err := sc.ReadString('\n')
	if err != nil || !isTerminated(input) {
		answer = INVALID_INPUT
	} else {
		input = StripInput(input)

		if isValidIP(input) {
			answer = GetAnswer(input)
		} else if input == "help" {
			answer = HELP_MESSAGE
		} else {
			answer = INVALID_INPUT
		}
	}

	conn.Write([]byte(answer))
}
