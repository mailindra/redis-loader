package service

import (
	"context"
	"fmt"
	"redis-loader/pkg/utils"
	"sync/atomic"
	"time"
)

const (
	batchSize = 1000 // Number of items per batch
	workers   = 8    // Number of parallel workers
)

type Repository interface {
	ParallelBatchSet(ctx context.Context, dataChan <-chan map[string]string, workers int) error
}

type LoaderService struct {
	repo Repository
}

func NewLoaderService(repo Repository) *LoaderService {
	return &LoaderService{
		repo: repo,
	}
}

func (s *LoaderService) LoadRandomData(ctx context.Context, count int) error {
	fmt.Printf("Loading %d key-value pairs into Redis...\n", count)
	startTime := time.Now()

	// Create channel for batches
	batchChan := make(chan map[string]string, workers*2) // Buffer the channel

	// Create counter for progress tracking
	var processed int64

	// Start progress reporter
	done := make(chan bool)
	go s.reportProgress(count, &processed, startTime, done)

	// Start generator goroutine
	go func() {
		defer close(batchChan)

		batch := make(map[string]string, batchSize)
		for i := 0; i < count; i++ {
			key := fmt.Sprintf("key:%s", utils.GenerateRandomString(8))
			value := utils.GenerateRandomString(16)
			batch[key] = value

			// When batch is full, send it and create new batch
			if len(batch) >= batchSize {
				batchChan <- batch
				// Update processed count for the batch
				atomic.AddInt64(&processed, int64(len(batch)))
				batch = make(map[string]string, batchSize)
			}
		}

		// Send remaining items
		if len(batch) > 0 {
			batchChan <- batch
			// Update processed count for the final batch
			atomic.AddInt64(&processed, int64(len(batch)))

		}
	}()

	// Process batches in parallel
	err := s.repo.ParallelBatchSet(ctx, batchChan, workers)

	// Update final count and wait for progress reporter to finish
	//atomic.AddInt64(&processed, int64(count))
	//<-done

	if err != nil {
		return fmt.Errorf("error loading data: %w", err)
	}

	// Wait for progress reporter to finish
	<-done

	elapsed := time.Since(startTime)
	fmt.Printf("\nData loading completed! Total time: %s\n", elapsed)
	fmt.Printf("Average speed: %.2f items/second\n", float64(count)/elapsed.Seconds())

	return nil
}

func (s *LoaderService) reportProgress(total int, processed *int64, startTime time.Time, done chan bool) {
	defer close(done)

	ticker := time.NewTicker(500 * time.Millisecond) // Update more frequently
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			current := atomic.LoadInt64(processed)
			if current >= int64(total) {
				return
			}

			elapsed := time.Since(startTime)
			speed := float64(current) / elapsed.Seconds()
			remaining := float64(total-int(current)) / speed

			fmt.Printf("\rProgress: %d/%d (%.2f%%) - Speed: %.2f items/sec - ETA: %s",
				current, total,
				float64(current)/float64(total)*100,
				speed,
				time.Duration(remaining*float64(time.Second)))
		}
	}
}
