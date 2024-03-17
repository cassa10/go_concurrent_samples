package high_order

import "fmt"

func Sample() {
	idX := composeFinal(id, id, "a")
	fmt.Println(fmt.Sprintf("idX: %v", idX))

	multiple2 := func(x int) int { return x * 2 }
	multiple4 := twice[int](multiple2)
	sample2 := multiple4(10)
	fmt.Println(fmt.Sprintf("sample2: %d", sample2))
}

// GenericFn :: (a -> b)
type GenericFn[T any] func(T) T

// id :: (a -> b)
func id(x any) any {
	return x
}

// twice :: (a -> b) -> (a -> b)
func twice[T any](fn GenericFn[T]) GenericFn[T] {
	return compose1(fn, fn)
}

// composeFinal :: ((b -> c), (a -> b), a) -> c
func composeFinal[T any](f GenericFn[T], g GenericFn[T], x T) T {
	return f(g(x))
}

// compose1 :: ((b -> c), (a -> b)) -> (a -> c)
func compose1[T any](f GenericFn[T], g GenericFn[T]) GenericFn[T] {
	return func(x T) T {
		return f(g(x))
	}
}

// compose2 :: (b -> c) -> ((a -> b) -> (a -> c))
func compose2[T any](f GenericFn[T]) func(GenericFn[T]) GenericFn[T] {
	return func(g GenericFn[T]) GenericFn[T] {
		return func(x T) T {
			return f(g(x))
		}
	}
}

// compose :: ((b -> c) -> ((a -> b) -> (a -> c)))
func compose[T any]() func(GenericFn[T]) func(GenericFn[T]) GenericFn[T] {
	return func(f GenericFn[T]) func(GenericFn[T]) GenericFn[T] {
		return func(g GenericFn[T]) GenericFn[T] {
			return func(x T) T {
				return f(g(x))
			}
		}
	}
}
