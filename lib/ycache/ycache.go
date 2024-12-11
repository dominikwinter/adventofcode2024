package ycache

import "sync"

type Func[I, O any] func(I) O
type TagFunc[I, O any] func(Func[I, O]) Func[I, O]
type CombinatorFunc[I, O any] func(CombinatorFunc[I, O]) Func[I, O]

func Y[I, O any](f TagFunc[I, O]) Func[I, O] {
	return func(self CombinatorFunc[I, O]) Func[I, O] {
		return f(func(n I) O {
			return self(self)(n)
		})
	}(func(self CombinatorFunc[I, O]) Func[I, O] {
		return f(func(n I) O {
			return self(self)(n)
		})
	})
}

func Cache[I comparable, O any](f TagFunc[I, O]) TagFunc[I, O] {
	var mu sync.RWMutex
	cache := make(map[I]O)

	return func(recurse Func[I, O]) Func[I, O] {
		return func(n I) O {
			mu.RLock()
			if ret, ok := cache[n]; ok {
				mu.RUnlock()
				return ret
			}
			mu.RUnlock()

			ret := f(recurse)(n)

			mu.Lock()
			cache[n] = ret
			mu.Unlock()

			return ret
		}
	}
}
