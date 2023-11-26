package chain

import "github.com/samber/lo"

func (c SliceChain[T]) Find(predicate func(item T) bool) (T, bool) {
	return lo.Find(c.value, predicate)
}

func (c SliceChain[T]) FindIndexOf(predicate func(item T) bool) (T, int, bool) {
	return lo.FindIndexOf(c.value, predicate)
}

func (c SliceChain[T]) FindLastIndexOf(predicate func(item T) bool) (T, int, bool) {
	return lo.FindLastIndexOf(c.value, predicate)
}

func (c SliceChain[T]) FindOrElse(fallback T, predicate func(item T) bool) T {
	return lo.FindOrElse(c.value, fallback, predicate)
}

func (c SliceChain[T]) MinBy(comparison func(a T, b T) bool) T {
	return lo.MinBy(c.value, comparison)
}

func (c SliceChain[T]) MaxBy(comparison func(a T, b T) bool) T {
	return lo.MaxBy(c.value, comparison)
}

func (c SliceChain[T]) Last() (T, error) {
	return lo.Last(c.value)
}

func (c SliceChain[T]) Sample() T {
	return lo.Sample(c.value)
}

func (c SliceChain[T]) Samples(count int) SliceChain[T] {
	return NewSliceChain(lo.Samples(c.value, count))
}

func (c ComparableSliceChain[T]) IndexOf(element T) int {
	return lo.IndexOf(c.value, element)
}

func (c ComparableSliceChain[T]) LastIndexOf(element T) int {
	return lo.LastIndexOf(c.value, element)
}

func (c ComparableSliceChain[T]) FindUniques() ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.FindUniques(c.value))
}

func (c ComparableSliceChain[T]) FindDuplicates() ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.FindDuplicates(c.value))
}

func (c OrderedSliceChain[T]) Min() T {
	return lo.Min(c.value)
}

func (c OrderedSliceChain[T]) Max() T {
	return lo.Max(c.value)
}

func (c MapChain[V, K]) FindKeyBy(predicate func(key K, value V) bool) (K, bool) {
	return lo.FindKeyBy(c.value, predicate)
}

func (c ComparableMapChain[V, K]) FindKey(value V) (K, bool) {
	return lo.FindKey(c.value, value)
}
