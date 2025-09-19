package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func asChan(ctx context.Context, vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			select {
			case c <- v:
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-ctx.Done():
				close(c)
				return
			}
		}
	}()
	return c
}

func merge(ctx context.Context, chans ...<-chan int) <-chan int {
	c := make(chan int)

	go func() {
		wg := sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case val, ok := <-ch:
						if !ok {
							return
						}
						select {
						case c <- val:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}()
		}
		wg.Wait()
		close(c)
	}()

	return c
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 1000)
	defer cancel()

	rand.Seed(time.Now().Unix())
	a := asChan(ctx, 1, 3, 5, 7)
	b := asChan(ctx, 2, 4, 6, 8)
	c := merge(ctx, a, b)
	for v := range c {
		fmt.Print(v)
	}

}
