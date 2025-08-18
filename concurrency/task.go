package concurrency

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Task[T any] func(ctx context.Context) (T, error)

func ConTasks[T any](ctx context.Context, tasks []Task[T], concurrency int) ([]T, error) {
	n := len(tasks)
	var mu sync.Mutex 
	results := make([]T, n)
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)
	for i, t := range tasks {
		idx, task := i, t
		g.Go(func() error {
			rst, err := task(ctx)
			if err != nil {
				return err
			}
			mu.Lock()
			results[idx] = rst
			mu.Unlock()
			return nil
		})
	}
	return results, g.Wait()
}
