package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type ErrorsCount struct {
	count int32
}

func (e *ErrorsCount) Get() int32 {
	return atomic.LoadInt32(&e.count)
}

func (e *ErrorsCount) Inc() {
	atomic.AddInt32(&e.count, 1)
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	ch := make(chan struct{}, n)
	errors := &ErrorsCount{}

	var wg sync.WaitGroup

	for i := 0; i < len(tasks); i++ {
		t := tasks[i]
		if errors.Get() >= int32(m) {
			break
		}

		wg.Add(1)

		go func(t Task) {
			defer wg.Done()

			if t() != nil {
				errors.Inc()
			}
			<-ch
		}(t)

		ch <- struct{}{}
	}

	wg.Wait()

	if errors.Get() > 0 {
		return ErrErrorsLimitExceeded
	}

	return nil
}
