package chain

import "github.com/samber/lo"

func ChunkString[T ~string](str T, size int) ComparableSliceChain[T] {
	return NewComparableSliceChain(lo.ChunkString(str, size))
}
