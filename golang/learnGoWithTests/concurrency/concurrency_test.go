package concurrency

import (
	"maps"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "philipkrueck.com"
}

func TestCheckWebsites(t *testing.T) {
	// given
	urls := []string{
		"google.com",
		"cloudflare.com",
		"philipkrueck.com",
		"x.com",
	}

	// when
	got := CheckWebsites(mockWebsiteChecker, urls)

	// then
	want := map[string]bool{
		"google.com":       true,
		"cloudflare.com":   true,
		"philipkrueck.com": false,
		"x.com":            true,
	}
	if !maps.Equal(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	// setup
	urls := make([]string, 100)

	// bench
	for b.Loop() {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
