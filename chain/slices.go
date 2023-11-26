package chain

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

type SliceChain[T any] struct {
	value []T
}

func NewSliceChain[T any](slice []T) SliceChain[T] {
	return SliceChain[T]{slice}
}

func Times[T any](count int, iteratee func(index int) T) SliceChain[T] {
	return NewSliceChain(lo.Times(count, iteratee))
}

func RepeatBy[T any](count int, predicate func(index int) T) SliceChain[T] {
	return NewSliceChain(lo.RepeatBy(count, predicate))
}

func (c SliceChain[T]) Value() []T {
	return c.value
}

func (c SliceChain[T]) Filter(predicate func(item T, index int) bool) SliceChain[T] {
	return NewSliceChain(lo.Filter(c.value, predicate))
}

func (c SliceChain[T]) ForEach(iteratee func(item T, index int)) SliceChain[T] {
	lo.ForEach(c.value, iteratee)
	return c
}

func (c SliceChain[T]) Interleave(collections ...[]T) SliceChain[T] {
	args := make([][]T, len(collections)+1)
	args = append(args, c.value)
	args = append(args, collections...)
	return NewSliceChain(lo.Interleave(args...))
}

func (c SliceChain[T]) Shuffle() SliceChain[T] {
	return NewSliceChain(lo.Shuffle(c.value))
}

func (c SliceChain[T]) Reverse() SliceChain[T] {
	return NewSliceChain(lo.Shuffle(c.value))
}

func (c SliceChain[T]) Drop(n int) SliceChain[T] {
	return NewSliceChain(lo.Drop(c.value, n))
}

func (c SliceChain[T]) DropRight(n int) SliceChain[T] {
	return NewSliceChain(lo.DropRight(c.value, n))
}

func (c SliceChain[T]) DropWhile(predicate func(item T) bool) SliceChain[T] {
	return NewSliceChain(lo.DropWhile(c.value, predicate))
}

func (c SliceChain[T]) DropRightWhile(predicate func(item T) bool) SliceChain[T] {
	return NewSliceChain(lo.DropRightWhile(c.value, predicate))
}

func (c SliceChain[T]) Reject(predicate func(item T, index int) bool) SliceChain[T] {
	return NewSliceChain(lo.Reject(c.value, predicate))
}

func (c SliceChain[T]) CountBy(predicate func(item T) bool) (count int) {
	return lo.CountBy(c.value, predicate)
}

func (c SliceChain[T]) Subset(offset int, length uint) SliceChain[T] {
	return NewSliceChain(lo.Subset(c.value, offset, length))
}

func (c SliceChain[T]) Slice(start int, end int) SliceChain[T] {
	return NewSliceChain(lo.Slice(c.value, start, end))
}

type ComparableSliceChain[T comparable] struct {
	SliceChain[T]
}

func NewComparableSliceChain[T comparable](slice []T) ComparableSliceChain[T] {
	base := NewSliceChain(slice)
	return ComparableSliceChain[T]{base}
}

func (c ComparableSliceChain[T]) Uniq() ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.Uniq(c.value))
}

func (c ComparableSliceChain[T]) Count(value T) (count int) {
	return lo.Count(c.value, value)
}

func (c ComparableSliceChain[T]) CountValues() MapChain[int, T] {
	return NewMapChain(lo.CountValues(c.value))
}

func (c ComparableSliceChain[T]) Replace(old T, new T, n int) ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.Replace(c.value, old, new, n))
}

func (c ComparableSliceChain[T]) ReplaceAll(old T, new T) ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.ReplaceAll(c.value, old, new))
}

func (c ComparableSliceChain[T]) Compact() ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.Compact(c.value))
}

type OrderedSliceChain[T constraints.Ordered] struct {
	ComparableSliceChain[T]
}

func NewOrderedSliceChain[T constraints.Ordered](slice []T) OrderedSliceChain[T] {
	base := NewComparableSliceChain(slice)
	return OrderedSliceChain[T]{base}
}

func (c OrderedSliceChain[T]) IsSorted() bool {
	return lo.IsSorted(c.value)
}
