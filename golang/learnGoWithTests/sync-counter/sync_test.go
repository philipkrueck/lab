package synccounter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times should result in 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		verifyCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		expectedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(expectedCount)

		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		verifyCounter(t, counter, expectedCount)
	})
}

func verifyCounter(t *testing.T, counter *Counter, expected int) {
	if counter.Value() != expected {
		t.Errorf("result: %d, expected: %d", counter.Value(), expected)
	}
}
