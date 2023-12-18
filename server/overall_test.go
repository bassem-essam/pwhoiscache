package server

import (
	"github.com/bassem-essam/pwhoiscache/client"
	"testing"
)

var originaClient *client.Client = client.New("whois.pwhois.org:43")

func TryOriginal(input string) string {
	return originaClient.Query(input)
}

func TryOurs(input string) string {
	return cli.Query(input)
}

func TestEqualityValid(t *testing.T) {
	ips := []string{
		"1.1.1.1",
		"1.1.1.2",
		"1.1.1.3",
		"8.8.8.8",
		"8.8.8.7",
		"8.8.8.100",
		"13.33.33.37",
		"13.33.33.37",
		"13.33.33.37",
		"13.33.13.37",
	}

	for _, ip := range ips {
		want := TryOriginal(ip)
		got := TryOurs(ip)

		if want != got {
			t.Errorf("Expected %s, got %s", want, got)
		}
	}
}

func TestEqualityEmpty(t *testing.T) {
	ips := []string{
		"0.0.0.0",
		"10.0.0.0",
	}

	for _, ip := range ips {
		want := TryOriginal(ip)
		got := TryOurs(ip)

		if want != got {
			t.Errorf("Expected %s, got %s", want, got)
		}
	}
}

func TestEqualityInvalid(t *testing.T) {
	ips := []string{
		"a.b.c.d",
		"1.2.3",
		"5000.100.33.25",
		"hello",
		// Notice that this is an invalid IP but the want server still returns an empty response
		// "300.300.300.300",
		// I think that also should not pass, but it passes anyway
		// "",
	}

	for _, ip := range ips {
		want := TryOriginal(ip)
		got := TryOurs(ip)

		if want != got {
			t.Errorf("%s Expected \n(%v)\n, got \n(%v)\n", ip, []byte(want), []byte(got))
		}
	}
}
