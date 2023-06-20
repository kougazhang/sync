package sync

import (
	"testing"
	"time"
)

func TestWaitGroup_Wait(t *testing.T) {
	max := NewWaitGroup(10)
	for i := 0; i < 10; i++ {
		max.Add()
		i := i
		go func() {
			defer max.Done()
			println(i)
			time.Sleep(10 * time.Millisecond)
		}()
	}
	max.Wait()
}
