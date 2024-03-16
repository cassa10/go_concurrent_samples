package channel_responses

import (
	"fmt"
	"log"
)

func Sample1() {
	logger := log.Default()
	ch1 := make(chan searchRes[SomeType])
	ch2 := make(chan searchRes[SomeType2])

	handleAsyncSearch1(logger, ch1, func() ([]SomeType, error) {
		return []SomeType{
			{Name: "Pepe", Age: 23},
			{Name: "Gerardo", Age: 63},
		}, fmt.Errorf("conection error")
	})
	handleAsyncSearch1(logger, ch2, func() ([]SomeType2, error) {
		delayInSeconds(10)
		return []SomeType2{
			{Name: "Pepe", Email: "pepe@a.c"},
		}, fmt.Errorf("unexpected error")
	})
	logger.Println(fmt.Sprintf("Waiting..."))
	res1 := <-ch1 //blocking until listen
	res2 := <-ch2 //blocking until listen
	logger.Println(fmt.Sprintf("Done! now ... processing"))

	if res1.Error != nil && res2.Error != nil {
		logger.Println("ALL ERROR!!!!")
	}

	logger.Println(fmt.Sprintf("Res1: %+v", res1))
	logger.Println(fmt.Sprintf("Res2: %+v", res2))
}

func handleAsyncSearch1[T any](logger *log.Logger, chRes chan searchRes[T], fn func() ([]T, error)) {
	go func(ch chan searchRes[T]) {
		res, err := fn()
		if err != nil {
			logger.Println(fmt.Errorf("some error found: %w", err))
		}
		ch <- searchRes[T]{Results: res, Error: err} //blocking when some ch listen
	}(chRes)
}
