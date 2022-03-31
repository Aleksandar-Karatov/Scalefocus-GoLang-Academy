package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := NewBufferedContext(55*time.Second, 1)

	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":

				time.Sleep(time.Second * 1)
				// try different values here
				fmt.Println("bar")

			}
		}
	})
}

type BufferedContext struct {
	context.Context /* Add other fields you might need */
	buffer          chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout) /*Implement the rest */
	buff := make(chan string, bufferSize)
	newBufferCTX := &BufferedContext{Context: ctx, buffer: buff, CancelFunc: cancel}
	//defer cancel()
	return newBufferCTX
}
func (bc *BufferedContext) Done() <-chan struct{} {
	/* This function will serve in place of the oriignal context */ /* make it so that the result channel gets closed in one of the to cases;
	      a) the emebdded context times out
		   b) the buffer gets filled    */

	//timeToEnd, ok := bc.Deadline()

	// err := bc.Err()
	// ch := make(chan struct{}, bc.buffSize)
	// go func(chan struct{}) {
	// 	for {
	// 		timeout := ok && time.Now().After(timeToEnd)
	// 		if timeout {

	// 			ch <- <-bc.Context.Done()
	// 			close(ch)
	// 			return
	// 		}

	// 		if err != nil {
	// 			ch <- <-bc.Context.Done()
	// 			close(ch)
	// 			return
	// 		}
	// 	}

	// }(ch)

	//return ch

	if len(bc.buffer) == cap(bc.buffer) {
		fmt.Println("Buffer is full")
		bc.CancelFunc()
	}
	return bc.Context.Done()
}
func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	/* This function serves for executing the test */ /* Implement the rest */

	fn(bc, bc.buffer)

}
