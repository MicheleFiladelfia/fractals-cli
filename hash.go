package main

import (
	"encoding/binary"
	"math"
)

const (
	fnvOffset64 = 14695981039346656037
	fnvPrime64  = 1099511628211
)

func hashParameters(cr, ci float64, maxIterations int, bailout float64) uint64 {
	var buf [28]byte
	binary.LittleEndian.PutUint64(buf[0:8], math.Float64bits(cr))
	binary.LittleEndian.PutUint64(buf[8:16], math.Float64bits(ci))
	binary.LittleEndian.PutUint32(buf[16:20], uint32(maxIterations))
	binary.LittleEndian.PutUint64(buf[20:28], math.Float64bits(bailout))

	var h uint64 = fnvOffset64
	for _, b := range buf {
		h ^= uint64(b)
		h *= fnvPrime64
	}
	return h
}