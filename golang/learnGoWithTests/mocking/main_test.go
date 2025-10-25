package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

type SpyCountdownOperations struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("correct output", func(t *testing.T) {
		// given
		buffer := &bytes.Buffer{}
		spySleeper := &SpyCountdownOperations{}

		// when
		Countdown(buffer, spySleeper)

		// then
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
	t.Run("correct order", func(t *testing.T) {
		// given
		spyCountdownOperations := &SpyCountdownOperations{}

		// when
		Countdown(spyCountdownOperations, spyCountdownOperations)

		// then
		got := spyCountdownOperations.Calls
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

type SpyTime struct {
	durationSleep time.Duration
}

func (st *SpyTime) SetSleepDuration(duration time.Duration) {
	st.durationSleep = duration
}

func TestConfigurableSleeper(t *testing.T) {
	// given
	sleepTime := time.Second * 5

	// when
	st := SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, st.SetSleepDuration}
	sleeper.Sleep()

	// then
	if st.durationSleep != sleepTime {
		t.Errorf("Slept for %v, but only slept for %v", sleepTime, st.durationSleep)
	}
}
