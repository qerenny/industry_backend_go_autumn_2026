package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if len(in) == 0 {
		return []Result[R]{}, nil
	}

	results := make([]Result[R], len(in))
	jobs := make(chan int)
	workers = min(workers, len(in))
	var wg sync.WaitGroup

	for range workers {
		wg.Go(func() {
			for i := range jobs {
				value, err := fn(ctx, in[i])
				var zero R

				if err != nil {
					results[i] = Result[R]{
						Value: zero,
						Err:   err,
					}
				} else {
					results[i] = Result[R]{
						Value: value,
						Err:   err,
					}
				}
			}
		})
	}
feed:
	for i := range in {
		if err := ctx.Err(); err != nil {
			break feed
		}
		select {
		case jobs <- i:
		case <-ctx.Done():
			break feed
		}
	}

	close(jobs)
	wg.Wait()

	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
