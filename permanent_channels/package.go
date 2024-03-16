package permanent_channels

import (
	"fmt"
	"log"
	"sync"
)

var wg = &sync.WaitGroup{}

const doneMSG = "DONE"

func samples() {
	logger := log.Default()
	ch := make(chan string, 1000)
	msgs := []string{
		"hola",
		"test",
		"test2",
		"chau",
	}
	nWorkers := 11
	for i := 1; i <= nWorkers; i++ {
		go writeMsgs(ch, msgs, i)
	}

	for {
		for msg := range ch {
			if msg == doneMSG {
				break
			} else {
				logger.Println(msg)
			}
		}
	}

	logger.Println("FINISHED OK")
}

func writeMsgs(ch chan string, msgs []string, nWorker int) {
	for i, msg := range msgs {
		//time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		select {
		case <-ch:
			return
		default:
			ch <- fmt.Sprintf("Worker%v-%v: %s", nWorker, i, msg)
		}
	}
	ch <- doneMSG
}
