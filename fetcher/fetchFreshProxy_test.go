package fetcher

import (
	"testing"
)

func TestFetchFreshProxy(t *testing.T) {
	if err := FetchFreshProxy(); err != nil {
		t.Fail()
	}
}
