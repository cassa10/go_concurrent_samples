package channels_simple

import (
	"fmt"
	"log"
)

func samples() {
	doAllIf(false, []func(){
		blockChannel,
		blockChannel2,
	})
	doAllIf(true, []func(){
		badNonBlockChannel,
		okNonBlockChannel,
	})
}

// DEADLOCK!!
func blockChannel() {
	logger := log.Default()
	msgCh := make(chan int) // blocks when channel is full

	msgCh <- 10 // op blocks thread because full channel

	msg := <-msgCh // never execute

	logWithMethod(logger, "blockChannel", msg)
}

// DEADLOCK!!
func blockChannel2() {
	logger := log.Default()
	msgCh := make(chan int, 2) // blocks when channel is full (2 messages)

	msgCh <- 1
	msgCh <- 3 // op blocks thread because full channel
	msgCh <- 10

	msg := <-msgCh // never execute

	logWithMethod(logger, "blockChannel2", msg)
}

func badNonBlockChannel() {
	logger := log.Default()
	msgCh := make(chan int, 2) // blocks when channel is full (2 messages)

	msgCh <- 1
	msgCh <- 3

	msg := <-msgCh // never execute

	logWithMethod(logger, "badNonBlockChannel", msg)
}

func okNonBlockChannel() {
	logger := log.Default()
	msgCh := make(chan int) // blocks when channel is full

	go func() {
		msgCh <- 10
	}()

	msg := <-msgCh // never execute

	logWithMethod(logger, "okNonBlockChannel", msg)
}

func doAllIf(b bool, fns []func()) {
	if b {
		for _, fn := range fns {
			fn()
		}
	}
}

func logWithMethod(logger *log.Logger, method, msg any) {
	logger.Println(fmt.Sprintf("(%s) %v", method, msg))
}
