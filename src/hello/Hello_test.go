package main

import (
	"testing"
)

func testHello(t *testing.T) {
	var expected string = "Hello"
	var actual string = Hello()

	if(actual != expected) {
		t.Fail()
	}
}