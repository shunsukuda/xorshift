package xrand

import "math/bits"

// https://prng.di.unimi.it/xoroshiro256plusplus.c

type Xoroshiro256pp struct {
	s [4]uint64
}

func (x *Xoroshiro256pp) Seed(seed int64) {
	s := NewSplitMix64(seed)
	x.s[0] = s.Uint64()
	x.s[1] = s.Uint64()
	x.s[2] = s.Uint64()
	x.s[3] = s.Uint64()
}

func (x *Xoroshiro256pp) Uint64() uint64 {
	result := bits.RotateLeft64(x.s[0]+x.s[3], 23) + x.s[0]
	t := x.s[1] << 17

	x.s[2] ^= x.s[0]
	x.s[3] ^= x.s[1]
	x.s[1] ^= x.s[2]
	x.s[0] ^= x.s[3]

	x.s[2] ^= t
	x.s[3] = bits.RotateLeft64(x.s[3], 45)

	return result
}

func (x *Xoroshiro256pp) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoroshiro256pp) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

// Call Uint64() * 2^128
func (x *Xoroshiro256pp) Jump() {
	var (
		jump = [...]uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}
	)

	s := [...]uint64{0, 0, 0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[0]
				s[1] ^= x.s[1]
				s[2] ^= x.s[2]
				s[3] ^= x.s[3]
			}
		}
		x.Uint64()
	}

	x.s[0] = s[0]
	x.s[1] = s[1]
	x.s[2] = s[2]
	x.s[3] = s[3]
}

// Call Uint64() * 2^192
func (x *Xoroshiro256pp) LongJump() {
	var (
		jump = [...]uint64{0x76e15d3efefdcbbf, 0xc5004e441c522fb3, 0x77710069854ee241, 0x39109bb02acbe635}
	)

	s := [...]uint64{0, 0, 0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[0]
				s[1] ^= x.s[1]
				s[2] ^= x.s[2]
				s[3] ^= x.s[3]
			}
		}
		x.Uint64()
	}

	x.s[0] = s[0]
	x.s[1] = s[1]
	x.s[2] = s[2]
	x.s[3] = s[3]
}
