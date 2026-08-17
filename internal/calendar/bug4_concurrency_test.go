package calendar

import (
	"sync"
	"testing"
	"time"
)

func TestBug4ConcurrentCalendarQueries(t *testing.T) {
	c, err := New(Config{Timezone: "Asia/Shanghai", Weekend: []int{6, 0}})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	const workers = 64
	var wg sync.WaitGroup
	start := make(chan struct{})
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-start
			for j := 0; j < 100; j++ {
				c.IsBusinessDay(Date{Y: 2026, M: time.August, D: 1 + (id+j)%28})
			}
		}(i)
	}
	close(start)
	wg.Wait()
}
