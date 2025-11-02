package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("make sure the fast server returns first", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)
		if err != nil {
			t.Errorf("Expected no error, but received: %v", err)
		}

		if got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	})
	t.Run("return an error if the server doesn't respond within 10 seconds", func(t *testing.T) {
		server := makeDelayedServer(11 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error, but got nil")
		}
	})
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}
