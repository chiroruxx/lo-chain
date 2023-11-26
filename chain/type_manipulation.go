package chain

import "github.com/samber/lo"

func FromAnySlice[T any](in []any) (out SliceChain[T], ok bool) {
	res, ok := lo.FromAnySlice[T](in)
	return NewSliceChain(res), ok
}

func (c SliceChain[T]) ToAnySlice() SliceChain[any] {
	return NewSliceChain(lo.ToAnySlice(c.value))
}
