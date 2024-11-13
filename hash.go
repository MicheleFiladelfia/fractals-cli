package main

import (
	"encoding/binary"
	"hash/fnv"
	"math"
)

func hashParameters(cr, ci float64, maxIterations int, bailout float64) uint64 {
	h := fnv.New64a()

	crBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(crBytes, math.Float64bits(cr))
	h.Write(crBytes)

	ciBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(ciBytes, math.Float64bits(ci))
	h.Write(ciBytes)

	maxIterationsBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(maxIterationsBytes, uint32(maxIterations))
	h.Write(maxIterationsBytes)

	bailoutBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bailoutBytes, math.Float64bits(bailout))
	h.Write(bailoutBytes)

	return h.Sum64()
}