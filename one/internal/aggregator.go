package internal

import (
	"context"
	"log"
	"sync"
	"time"
)

// Web Client interface
type Client interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

// Worker struct
type worker struct {
	client Client
}

func Run(c Client, workerNum int) int64 {
	wg := &sync.WaitGroup{}

	byteCh := make(chan []byte, workerNum) // create buffered channel that can keep all fetchec info

	for i := 0; i < workerNum; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()                                  // marking that worker finiched work
			worker{c}.fetchURL("https://google.com", byteCh) // launching worker
		}()
	}

	wg.Wait()     // waiting for all workers
	close(byteCh) // closing channel as no one will be writing to it

	countCh := make(chan int64)
	go aggregate(byteCh, countCh) // aggregating, could start right away and not wait for byteCh to be closed, but this was a requirement

	return <-countCh
}

func aggregate(byteCh <-chan []byte, countCh chan<- int64) {
	var count int

	for b := range byteCh {
		count = count + len(b)
	}

	countCh <- int64(count)
}

func (w worker) fetchURL(url string, byteCh chan<- []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	b, err := w.client.Get(ctx, url)
	if err != nil {
		log.Printf("worker failed, %v \n", err)
		return
	}

	byteCh <- b
}
