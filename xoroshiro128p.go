package xrand

import (
	"math"
	"math/bits"
)

// https://prng.di.unimi.it/xoroshiro128plus.c

type Xoroshiro128p struct {
	s [2]uint64
}

func NewXoroshiro128p(seed int64) *Xoroshiro128p {
	x := Xoroshiro128p{}
	x.Seed(seed)
	return &x
}

func (x *Xoroshiro128p) Seed(seed int64) {
	s := NewSplitMix64(seed)
	x.s[0] = s.Uint64()
	x.s[1] = s.Uint64()
}

func (x *Xoroshiro128p) Uint64() uint64 {
	s := [...]uint64{x.s[0], x.s[1]}
	result := s[0] + s[1]
	s[1] ^= s[0]
	x.s[0] = bits.RotateLeft64(s[0], 24) ^ s[1] ^ (s[1] << 16)
	x.s[1] = bits.RotateLeft64(s[1], 37)
	return result
}

func (x *Xoroshiro128p) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoroshiro128p) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

func (x *Xoroshiro128p) Float64() float64 {
	return math.Float64frombits(0x3ff<<52|x.Uint64()>>12) - 1.0
}

// Call Uint64() * 2^64
func (x *Xoroshiro128p) Jump() {
	var (
		jump = [...]uint64{0xdf900294d8f554a5, 0x170865df4b3201fc}
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
func (x *Xoroshiro128p) LongJump() {
	var (
		jump = [...]uint64{0xd2a98b26625eee7b, 0xdddf9b1090aa7ac1}
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
