package xrand

import (
	"math"
	"math/bits"
)

// https://prng.di.unimi.it/xoroshiro128plusplus.c

type Xoroshiro128pp struct {
	s [2]uint64
}

func NewXoroshiro128pp(seed int64) *Xoroshiro128pp {
	x := Xoroshiro128pp{}
	x.Seed(seed)
	return &x
}

func (x *Xoroshiro128pp) Seed(seed int64) {
	s := NewSplitMix64(seed)
	x.s[0] = s.Uint64()
	x.s[1] = s.Uint64()
}

func (x *Xoroshiro128pp) Uint64() uint64 {
	s := [...]uint64{x.s[0], x.s[1]}
	result := bits.RotateLeft64(s[0]+s[1], 17) + s[0]
	s[1] ^= s[0]
	x.s[0] = bits.RotateLeft64(s[0], 49) ^ s[1] ^ (s[1] << 21)
	x.s[1] = bits.RotateLeft64(s[1], 28)
	return result
}

func (x *Xoroshiro128pp) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoroshiro128pp) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

func (x *Xoroshiro128pp) Float64() float64 {
	return math.Float64frombits(0x3ff<<52|x.Uint64()>>12) - 1.0
}

// Call Uint64() * 2^64
func (x *Xoroshiro128pp) Jump() {
	var (
		jump = [...]uint64{0x2bd7a6a6e99c2ddc, 0x0992ccaf6a6fca05}
	)

	s := [...]uint64{0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[0]
				s[1] ^= x.s[1]
			}
		}
		x.Uint64()
	}

	x.s[0] = s[0]
	x.s[1] = s[1]
}

// Call Uint64() * 2^96
func (x *Xoroshiro128pp) LongJump() {
	var (
		jump = [...]uint64{0x360fd5f2cf8d5d99, 0x9c6e6877736c46e3}
	)

	s := [...]uint64{0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[0]
				s[1] ^= x.s[1]
			}
		}
		x.Uint64()
	}

	x.s[0] = s[0]
	x.s[1] = s[1]
}
