package chain

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

func Range(start int) OrderedSliceChain[int] {
	return NewOrderedSliceChain(lo.Range(start))
}

func RangeFrom[T constraints.Integer | constraints.Float](start T, elementNum int) OrderedSliceChain[T] {
	return NewOrderedSliceChain(lo.RangeFrom(start, elementNum))
}

func RangeWithSteps[T constraints.Integer | constraints.Float](start, end, step T) OrderedSliceChain[T] {
	return NewOrderedSliceChain(lo.RangeWithSteps(start, end, step))
}
