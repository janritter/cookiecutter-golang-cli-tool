package main

import (
	"testing"
)

func TestReturnGreeting(t *testing.T) {
	name := "test-name"
	greeting := returnGreeting(name)

	if greeting != "Hello World from test-name" {
		t.Errorf("Wrong greeting, was %s", greeting)
	}
}
