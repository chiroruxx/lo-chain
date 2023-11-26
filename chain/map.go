package chain

import "github.com/samber/lo"

type MapChain[V any, K comparable] struct {
	value map[K]V
}

func NewMapChain[T any, U comparable](m map[U]T) MapChain[T, U] {
	return MapChain[T, U]{m}
}

func FromEntries[V any, K comparable](entries []lo.Entry[K, V]) MapChain[V, K] {
	return NewMapChain(lo.FromEntries(entries))
}

func FromPairs[V any, K comparable](entries []lo.Entry[K, V]) MapChain[V, K] {
	return NewMapChain(lo.FromPairs(entries))
}

func (c MapChain[V, K]) Value() map[K]V {
	return c.value
}

func (c MapChain[V, K]) Keys() ComparableSliceChain[K] {
	return NewComparableSliceChain(lo.Keys(c.value))
}

func (c MapChain[V, K]) Values() SliceChain[V] {
	return NewSliceChain(lo.Values(c.value))
}

func (c MapChain[V, K]) ValurOr(key K, fallback V) V {
	return lo.ValueOr(c.value, key, fallback)
}

func (c MapChain[V, K]) PickBy(predicate func(key K, value V) bool) MapChain[V, K] {
	return NewMapChain(lo.PickBy(c.value, predicate))
}

func (c MapChain[V, K]) PickByKeys(keys []K) MapChain[V, K] {
	return NewMapChain(lo.PickByKeys(c.value, keys))
}

func (c MapChain[V, K]) OmitBy(predicate func(key K, value V) bool) MapChain[V, K] {
	return NewMapChain(lo.OmitBy(c.value, predicate))
}

func (c MapChain[V, K]) OmitByKeys(keys []K) MapChain[V, K] {
	return NewMapChain(lo.OmitByKeys(c.value, keys))
}

func (c MapChain[V, K]) Entries() SliceChain[lo.Entry[K, V]] {
	return NewSliceChain(lo.Entries(c.value))
}

func (c MapChain[V, K]) ToPairs() SliceChain[lo.Entry[K, V]] {
	return NewSliceChain(lo.ToPairs(c.value))
}

func (c MapChain[V, K]) Assign(maps ...map[K]V) MapChain[V, K] {
	args := make([]map[K]V, len(maps)+1)
	args = append(args, c.value)
	args = append(args, maps...)

	return NewMapChain(lo.Assign(args...))
}

type ComparableMapChain[V comparable, K comparable] struct {
	MapChain[V, K]
}

func NewComparableMapChain[V comparable, K comparable](m map[K]V) ComparableMapChain[V, K] {
	base := NewMapChain(m)
	return ComparableMapChain[V, K]{base}
}

func (c ComparableMapChain[V, K]) PickByValues(values []V) ComparableMapChain[V, K] {
	return NewComparableMapChain(lo.PickByValues(c.value, values))
}

func (c ComparableMapChain[V, K]) OmitByValues(values []V) ComparableMapChain[V, K] {
	return NewComparableMapChain(lo.OmitByValues(c.value, values))
}

func (c ComparableMapChain[V, K]) Invert() ComparableMapChain[K, V] {
	return NewComparableMapChain(lo.Invert(c.value))
}
