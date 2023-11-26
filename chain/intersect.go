package chain

import "github.com/samber/lo"

func (c SliceChain[T]) ContainsBy(predicate func(item T) bool) bool {
	return lo.ContainsBy(c.value, predicate)
}

func (c SliceChain[T]) EveryBy(predicate func(item T) bool) bool {
	return lo.EveryBy(c.value, predicate)
}

func (c SliceChain[T]) SomeBy(predicate func(item T) bool) bool {
	return lo.SomeBy(c.value, predicate)
}

func (c SliceChain[T]) NoneBy(predicate func(item T) bool) bool {
	return lo.NoneBy(c.value, predicate)
}

func (c ComparableSliceChain[T]) Contains(element T) bool {
	return lo.Contains(c.value, element)
}

func (c ComparableSliceChain[T]) Every(subset []T) bool {
	return lo.Every(c.value, subset)
}

func (c ComparableSliceChain[T]) Some(subset []T) bool {
	return lo.Some(c.value, subset)
}

func (c ComparableSliceChain[T]) None(subset []T) bool {
	return lo.None(c.value, subset)
}

func (c ComparableSliceChain[T]) Intersect(list []T) ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.Intersect(c.value, list))
}

func (c ComparableSliceChain[T]) Difference(list []T) (ComparableSliceChain[T], ComparableSliceChain[T]) {
	left, right := lo.Difference(c.value, list)
	return NewComparableSliceChain(left), NewComparableSliceChain(right)
}

func (c ComparableSliceChain[T]) Union(lists ...[]T) ComparableSliceChain[T] {
	args := make([][]T, len(lists)+1)
	args = append(args, c.value)
	args = append(args, lists...)
	return NewComparableSliceChain(lo.Union(args...))
}

func (c ComparableSliceChain[T]) Without(exclude ...T) ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.Without(c.value, exclude...))
}

func (c ComparableSliceChain[T]) WithoutEmpty() ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.WithoutEmpty(c.value))
}
