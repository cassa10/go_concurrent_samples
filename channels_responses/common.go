package channel_responses

import "time"

func delayInSeconds(sec int) {
	time.Sleep(time.Second * time.Duration(sec))
}

type searchRes[T any] struct {
	Results []T
	Error   error
}

type SomeType struct {
	Name string
	Age  int
}

type SomeType2 struct {
	Name  string
	Email string
}
