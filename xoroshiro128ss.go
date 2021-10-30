package xrand

import "math/bits"

// https://prng.di.unimi.it/xoroshiro128starstar.c

type Xoroshiro128ss struct {
	s [2]uint64
}

func (x *Xoroshiro128ss) Seed(seed int64) {
	s := NewSplitMix64(seed)
	x.s[0] = s.Uint64()
	x.s[1] = s.Uint64()
}

func (x *Xoroshiro128ss) Uint64() uint64 {
	s := [...]uint64{x.s[0], x.s[1]}
	result := bits.RotateLeft64(s[0]*5, 7) * 9
	s[1] ^= s[0]
	x.s[0] = bits.RotateLeft64(s[0], 24) ^ s[1] ^ (s[1] << 16)
	x.s[1] = bits.RotateLeft64(s[1], 37)
	return result
}

func (x *Xoroshiro128ss) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoroshiro128ss) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

// Call Uint64() * 2^64
func (x *Xoroshiro128ss) Jump() {
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
func (x *Xoroshiro128ss) LongJump() {
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
