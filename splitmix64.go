package xrand

import "fmt"

type SplitMix64 struct {
	s uint64
}

// https://prng.di.unimi.it/splitmix64.c

func NewSplitMix64(seed int64) SplitMix64 {
	var x SplitMix64
	x.Seed(seed)
	return x
}

func (x SplitMix64) State() uint64 {
	return uint64(x.s)
}

func (x *SplitMix64) Seed(seed int64) {
	x.s = unsafeInt64ToUint64(seed)
}

func (x *SplitMix64) Uint64() uint64 {
	x.s += 0x9e3779b97f4a7c15
	z := x.State()
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func (x *SplitMix64) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *SplitMix64) Int63() int64 {
	return int64(x.Uint64()) & (1<<63 - 1)
}

func (x SplitMix64) String() string {
	return fmt.Sprintf("%032x", x.State())
}

func (x SplitMix64) GoString() string {
	return "xrand.SplitMix64{state:\"" + x.String() + "\"}"
}
