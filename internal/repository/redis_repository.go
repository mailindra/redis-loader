package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

// BatchSet uses pipelining to set multiple key-value pairs efficiently
func (r *RedisRepository) BatchSet(ctx context.Context, pairs map[string]string) error {
	pipe := r.client.Pipeline()

	for k, v := range pairs {
		pipe.Set(ctx, k, v, 0)
	}

	_, err := pipe.Exec(ctx)
	return err
}

// ParallelBatchSet processes multiple batches in parallel using goroutines
func (r *RedisRepository) ParallelBatchSet(ctx context.Context, dataChan <-chan map[string]string, workers int) error {
	var wg sync.WaitGroup
	errorChan := make(chan error, workers)

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for batch := range dataChan {
				if err := r.BatchSet(ctx, batch); err != nil {
					errorChan <- err
					return
				}
			}
		}()
	}

	// Wait for all workers to finish
	wg.Wait()
	close(errorChan)

	// Check for any errors
	for err := range errorChan {
		if err != nil {
			return err
		}
	}

	return nil
}
