package xrand

import "math/bits"

// https://prng.di.unimi.it/xoshiro512starstar.c

type Xoshiro512ss struct {
	s [8]uint64
}

func (x *Xoshiro512ss) Seed(seed int64) {
	s := NewSplitMix64(seed)
	x.s[0] = s.Uint64()
	x.s[1] = s.Uint64()
	x.s[2] = s.Uint64()
	x.s[3] = s.Uint64()
	x.s[4] = s.Uint64()
	x.s[5] = s.Uint64()
	x.s[6] = s.Uint64()
	x.s[7] = s.Uint64()
}

func (x *Xoshiro512ss) Uint64() uint64 {
	result := bits.RotateLeft64(x.s[1]*5, 7) * 9
	t := x.s[1] << 11

	x.s[2] ^= x.s[0]
	x.s[5] ^= x.s[1]
	x.s[1] ^= x.s[2]
	x.s[7] ^= x.s[3]
	x.s[3] ^= x.s[4]
	x.s[4] ^= x.s[5]
	x.s[0] ^= x.s[6]
	x.s[6] ^= x.s[7]

	x.s[6] ^= t
	x.s[7] = bits.RotateLeft64(x.s[7], 21)

	return result
}

func (x *Xoshiro512ss) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoshiro512ss) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

// Call Uint64() * 2^256
func (x *Xoshiro512ss) Jump() {
	var (
		jump = [...]uint64{
			0x33ed89b6e7a353f9, 0x760083d7955323be, 0x2837f2fbb5f22fae, 0x4b8c5674d309511c,
			0xb11ac47a7ba28c25, 0xf1be7667092bcc1c, 0x53851efdb6df0aaf, 0x1ebbc8b23eaf25db,
		}
	)

	s := [...]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[0]
				s[1] ^= x.s[1]
				s[2] ^= x.s[2]
				s[3] ^= x.s[3]
				s[4] ^= x.s[4]
				s[5] ^= x.s[5]
				s[6] ^= x.s[6]
				s[7] ^= x.s[7]
			}
		}
		x.Uint64()
	}

	s[0] = x.s[0]
	s[1] = x.s[1]
	s[2] = x.s[2]
	s[3] = x.s[3]
	s[4] = x.s[4]
	s[5] = x.s[5]
	s[6] = x.s[6]
	s[7] = x.s[7]
}

// Call Uint64() * 2^192
func (x *Xoshiro512ss) LongJump() {
	var (
		jump = [...]uint64{
			0x11467fef8f921d28, 0xa2a819f2e79c8ea8, 0xa8299fc284b3959a, 0xb4d347340ca63ee1,
			0x1cb0940bedbff6ce, 0xd956c5c4fa1f8e17, 0x915e38fd4eda93bc, 0x5b3ccdfa5d7daca5,
		}
	)

	s := [...]uint64{0, 0, 0, 0, 0, 0, 0, 0}
	for i := range jump {
		for b := 0; b < 64; b++ {
			if (jump[i] & 1 << b) != 0 {
				s[0] ^= x.s[0]
				s[1] ^= x.s[1]
				s[2] ^= x.s[2]
				s[3] ^= x.s[3]
				s[4] ^= x.s[4]
				s[5] ^= x.s[5]
				s[6] ^= x.s[6]
				s[7] ^= x.s[7]
			}
		}
		x.Uint64()
	}

	s[0] = x.s[0]
	s[1] = x.s[1]
	s[2] = x.s[2]
	s[3] = x.s[3]
	s[4] = x.s[4]
	s[5] = x.s[5]
	s[6] = x.s[6]
	s[7] = x.s[7]
}
