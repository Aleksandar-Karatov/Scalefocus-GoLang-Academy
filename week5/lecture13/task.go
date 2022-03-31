package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := NewBufferedContext(time.Second, 10)

	ctx.Run(func(ctx context.Context, buffer chan string) {

		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 1)
				// try different values here
				fmt.Println("bar")

			}
		}
	})
}

type BufferedContext struct {
	context.Context /* Add other fields you might need */
	buffSize        int
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout) /*Implement the rest */
	var bc BufferedContext
	bc.buffSize = bufferSize
	bc.Context = ctx
	defer cancel()
	return &bc
}
func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */ /* make it so that the result channel gets closed in one of the to cases;
	      a) the emebdded context times out
		   b) the buffer gets filled    */
	ch := make(chan struct{})
	timeToEnd, _ := bc.Deadline()
	go func() {
		if bc.buffSize == 0 {
			ch <- <-bc.Context.Done()
			close(ch)
			return
		}
		if time.Now().Equal(timeToEnd) {
			ch <- <-bc.Context.Done()
			close(ch)
			return
		}
	}()
	return ch
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */ /* Implement the rest */

}
