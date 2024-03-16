package channel_responses

import (
	"fmt"
	"log"
	"sync"
)

func Sample2() {
	logger := log.Default()

	asyncCalls := 2
	ch := make(chan any, asyncCalls)
	wg := &sync.WaitGroup{}
	wg.Add(asyncCalls)

	handleAsyncSearch2(logger, wg, ch, func() ([]SomeType, error) {
		return []SomeType{
			{Name: "Pepe", Age: 23},
			{Name: "Gerardo", Age: 63},
		}, fmt.Errorf("some error sarasa")
	})
	handleAsyncSearch2(logger, wg, ch, func() ([]SomeType2, error) {
		delayInSeconds(3)
		return []SomeType2{
			{Name: "Pepe", Email: "pepe@a.c"},
		}, nil
	})

	logger.Println(fmt.Sprintf("Waiting..."))
	wg.Wait()
	close(ch) //IMPORTANT for no deadlock!!

	logger.Println(fmt.Sprintf("Done! now ... processing"))
	var respLs []interface{}
	for chResp := range ch {
		respLs = append(respLs, chResp)
	}
	logger.Println(fmt.Sprintf("RespLs: %+v", respLs))
	var res1 searchRes[SomeType]
	var res2 searchRes[SomeType2]
	for _, res := range respLs {
		switch res.(type) {
		case searchRes[SomeType]:
			res1 = res.(searchRes[SomeType])
		case searchRes[SomeType2]:
			res2 = res.(searchRes[SomeType2])
		default:
			logger.Println("error!!!!")
		}
	}

	logger.Println(fmt.Sprintf("Res1: %+v", res1))
	logger.Println(fmt.Sprintf("Res2: %+v", res2))
}

func handleAsyncSearch2[T any](logger *log.Logger, waitGroup *sync.WaitGroup, chRes chan any, fn func() ([]T, error)) {
	go func(ch chan any, wg *sync.WaitGroup) {
		res, err := fn()
		if err != nil {
			logger.Println(fmt.Errorf("some error found: %w", err))
		}
		ch <- searchRes[T]{Results: res, Error: err}
		logger.Println("DONE!")
		wg.Done()
	}(chRes, waitGroup)
}

func anyError(ls []searchRes[any]) bool {
	for _, elem := range ls {
		if elem.Error != nil {
			return true
		}
	}
	return false
}

func getError(resp any) error {
	r, ok := resp.(searchRes[any])
	if !ok {
		return fmt.Errorf("cannot get error")
	}
	return r.Error
}
