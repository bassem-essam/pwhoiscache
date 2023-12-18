package server

import (
	"fmt"
	"os"
	"github.com/bassem-essam/pwhoiscache/client"
	"testing"
	"time"
)

var cli *client.Client = client.New("localhost:43")

func TryInput(input string) string {
	return cli.Query(input)
}

func TestMain(m *testing.M) {
	fmt.Println("Setting up server...")
	go StartServer()
	fmt.Println("Server started!")
	time.Sleep(1 * time.Second)
	fmt.Println("Running tests...")
	result := m.Run()
	os.Exit(result)
}

func TestValidInput(t *testing.T) {
	answer := TryInput("1.2.3.4")
	if answer == INVALID_INPUT {
		t.Errorf("Expected valid answer, got %s", answer)
	}

	answer = TryInput("255.255.255.255")
	if answer == INVALID_INPUT {
		t.Errorf("Expected valid answer, got %s", answer)
	}
}

func TestHelp(t *testing.T) {
	answer := TryInput("help")
	if answer != HELP_MESSAGE {
		t.Errorf("Expected help message, got %s", answer)
	}
}

func TestEmptyResponse(t *testing.T) {
	answer := TryInput("10.10.10.10")
	if answer != "" {
		t.Errorf("Expected empty answer, got %s", answer)
	}

	answer = TryInput("0.0.0.0")
	if answer != "" {
		t.Errorf("Expected empty answer, got %s", answer)
	}
}

func TestInvalidInput(t *testing.T) {
	non_ips := []string{
		"1.2.3",
		"hello",
		"13333337",
		"a.b.c.d",
	}

	for _, input := range non_ips {
		if isValidIP(input) {
			t.Errorf("Expected %s to be invalid", input)
		}
	}

	bad_ips := []string{
		"999.999.999.999",
		"1.1.1.256",
		"-1.-1.-1.-1",
	}

	for _, input := range bad_ips {
		if isValidIP(input) {
			t.Errorf("Expected %s to be invalid", input)
		}
	}
}
