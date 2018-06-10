package main

import (
	"testing"
)

func TestFetchFreshProxy(t *testing.T) {
	if err := fetchFreshProxy(); err != nil {
		t.Fail()
	}
}
