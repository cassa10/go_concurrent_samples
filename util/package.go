package util

import "time"

func DelayInSeconds(sec int) {
	time.Sleep(time.Second * time.Duration(sec))
}

type SimpleRes[T any] struct {
	Result T
	Error  error
}

type SearchRes[T any] struct {
	Results []T
	Error   error
}

func DoAllIf(b bool, fns []func()) {
	if b {
		for _, fn := range fns {
			fn()
		}
	}
}
