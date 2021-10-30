package xrand

import "unsafe"

type SplitMix64 uint64

func NewSplitMix64(seed int64) SplitMix64 {
	var x SplitMix64
	x.Seed(seed)
	return x
}

func (x *SplitMix64) Seed(seed int64) {
	*x = SplitMix64(unsafeInt64ToUint64(seed))
}

func (x *SplitMix64) Uint64() uint64 {
	*x = *x + SplitMix64(0x9E3779B97F4A7C15)
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

func (x *SplitMix64) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *SplitMix64) Int63() int64 {
	return int64(x.Uint64()) & (1<<63 - 1)
}

func unsafeInt64ToUint64(v int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

func unsafeUint64ToInt64(v uint64) int64 {
	return *(*int64)(unsafe.Pointer(&v))
}
