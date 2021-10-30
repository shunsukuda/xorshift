package xrand

import "math/bits"

// https://prng.di.unimi.it/xoshiro1024starstar.c

type Xoshiro1024ss struct {
	s [16]uint64
	p int
}

func (x *Xoshiro1024ss) Seed(seed int64) {
	s := NewSplitMix64(seed)
	x.s[0] = s.Uint64()
	x.s[1] = s.Uint64()
	x.s[2] = s.Uint64()
	x.s[3] = s.Uint64()
	x.s[4] = s.Uint64()
	x.s[5] = s.Uint64()
	x.s[6] = s.Uint64()
	x.s[7] = s.Uint64()
	x.s[8] = s.Uint64()
	x.s[9] = s.Uint64()
	x.s[10] = s.Uint64()
	x.s[11] = s.Uint64()
	x.s[12] = s.Uint64()
	x.s[13] = s.Uint64()
	x.s[14] = s.Uint64()
	x.s[15] = s.Uint64()
}

func (x *Xoshiro1024ss) Uint64() uint64 {
	q := x.p
	x.p = (x.p + 1) & 15
	s0 := x.s[x.p]
	s15 := x.s[q]
	result := bits.RotateLeft64(s0*5, 7) * 9

	s15 ^= s0
	x.s[q] = bits.RotateLeft64(s0, 25) ^ s15 ^ (s15 << 27)
	x.s[x.p] = bits.RotateLeft64(s15, 36)

	return result
}

func (x *Xoshiro1024ss) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoshiro1024ss) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

// Call Uint64() * 2^512
func (x *Xoshiro1024ss) Jump() {
	var (
		jump = [...]uint64{
			0x931197d8e3177f17, 0xb59422e0b9138c5f, 0xf06a6afb49d668bb, 0xacb8a6412c8a1401,
			0x12304ec85f0b3468, 0xb7dfe7079209891e, 0x405b7eec77d9eb14, 0x34ead68280c44e4a,
			0xe0e4ba3e0ac9e366, 0x8f46eda8348905b7, 0x328bf4dbad90d6ff, 0xc8fd6fb31c9effc3,
			0xe899d452d4b67652, 0x45f387286ade3205, 0x03864f454a8920bd, 0xa68fa28725b1b384,
		}
	)

	s := [...]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[(0+x.p)&15]
				s[1] ^= x.s[(1+x.p)&15]
				s[2] ^= x.s[(2+x.p)&15]
				s[3] ^= x.s[(3+x.p)&15]
				s[4] ^= x.s[(4+x.p)&15]
				s[5] ^= x.s[(5+x.p)&15]
				s[6] ^= x.s[(6+x.p)&15]
				s[7] ^= x.s[(7+x.p)&15]
				s[8] ^= x.s[(8+x.p)&15]
				s[9] ^= x.s[(9+x.p)&15]
				s[10] ^= x.s[(10+x.p)&15]
				s[11] ^= x.s[(11+x.p)&15]
				s[12] ^= x.s[(12+x.p)&15]
				s[13] ^= x.s[(13+x.p)&15]
				s[14] ^= x.s[(14+x.p)&15]
				s[15] ^= x.s[(15+x.p)&15]
			}
		}
		x.Uint64()
	}

	x.s[(0+x.p)&15] = s[0]
	x.s[(1+x.p)&15] = s[1]
	x.s[(2+x.p)&15] = s[2]
	x.s[(3+x.p)&15] = s[3]
	x.s[(4+x.p)&15] = s[4]
	x.s[(5+x.p)&15] = s[5]
	x.s[(6+x.p)&15] = s[6]
	x.s[(7+x.p)&15] = s[7]
	x.s[(8+x.p)&15] = s[8]
	x.s[(9+x.p)&15] = s[9]
	x.s[(10+x.p)&15] = s[10]
	x.s[(11+x.p)&15] = s[11]
	x.s[(12+x.p)&15] = s[12]
	x.s[(13+x.p)&15] = s[13]
	x.s[(14+x.p)&15] = s[14]
	x.s[(15+x.p)&15] = s[15]
}

// Call Uint64() * 2^768
func (x *Xoshiro1024ss) LongJump() {
	var (
		jump = [...]uint64{
			0x7374156360bbf00f, 0x4630c2efa3b3c1f6, 0x6654183a892786b1, 0x94f7bfcbfb0f1661,
			0x27d8243d3d13eb2d, 0x9701730f3dfb300f, 0x2f293baae6f604ad, 0xa661831cb60cd8b6,
			0x68280c77d9fe008c, 0x50554160f5ba9459, 0x2fc20b17ec7b2a9a, 0x49189bbdc8ec9f8f,
			0x92a65bca41852cc1, 0xf46820dd0509c12a, 0x52b00c35fbf92185, 0x1e5b3b7f589e03c1,
		}
	)

	s := [...]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[(0+x.p)&15]
				s[1] ^= x.s[(1+x.p)&15]
				s[2] ^= x.s[(2+x.p)&15]
				s[3] ^= x.s[(3+x.p)&15]
				s[4] ^= x.s[(4+x.p)&15]
				s[5] ^= x.s[(5+x.p)&15]
				s[6] ^= x.s[(6+x.p)&15]
				s[7] ^= x.s[(7+x.p)&15]
				s[8] ^= x.s[(8+x.p)&15]
				s[9] ^= x.s[(9+x.p)&15]
				s[10] ^= x.s[(10+x.p)&15]
				s[11] ^= x.s[(11+x.p)&15]
				s[12] ^= x.s[(12+x.p)&15]
				s[13] ^= x.s[(13+x.p)&15]
				s[14] ^= x.s[(14+x.p)&15]
				s[15] ^= x.s[(15+x.p)&15]
			}
		}
		x.Uint64()
	}

	x.s[(0+x.p)&15] = s[0]
	x.s[(1+x.p)&15] = s[1]
	x.s[(2+x.p)&15] = s[2]
	x.s[(3+x.p)&15] = s[3]
	x.s[(4+x.p)&15] = s[4]
	x.s[(5+x.p)&15] = s[5]
	x.s[(6+x.p)&15] = s[6]
	x.s[(7+x.p)&15] = s[7]
	x.s[(8+x.p)&15] = s[8]
	x.s[(9+x.p)&15] = s[9]
	x.s[(10+x.p)&15] = s[10]
	x.s[(11+x.p)&15] = s[11]
	x.s[(12+x.p)&15] = s[12]
	x.s[(13+x.p)&15] = s[13]
	x.s[(14+x.p)&15] = s[14]
	x.s[(15+x.p)&15] = s[15]
}
