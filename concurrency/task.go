package concurrency

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Task[T any] func(ctx context.Context) (T, error)

func ConTasks[T any](ctx context.Context, tasks []Task[T], concurrency int) ([]*T, error) {
	n := len(tasks)
	resultPtrs := make([]*T, n)
	for i := 0; i < n; i++ {
		resultPtrs[i] = new(T)
	}
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(concurrency)
	for i, t := range tasks {
		idx, task := i, t // 确保协程参数是正确的
		rstPtr := resultPtrs[idx]
		g.Go(func() error {
			rst, err := task(ctx)
			if err != nil {
				return err
			}
			*rstPtr = rst
			return nil
		})
	}
	return resultPtrs, g.Wait()
}
