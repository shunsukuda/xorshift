package xrand

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/bits"
	"time"
)

// https://prng.di.unimi.it/xoshiro512plus.c

type Xoshiro512p struct {
	s [8]uint64
}

func NewXoshiro512p(seed int64) *Xoshiro512p {
	x := Xoshiro512p{}
	x.Seed(seed)
	return &x
}

func (x Xoshiro512p) State() []byte {
	s := make([]byte, 64)
	binary.BigEndian.PutUint64(s[:8], x.s[0])
	binary.BigEndian.PutUint64(s[8:16], x.s[1])
	binary.BigEndian.PutUint64(s[16:24], x.s[2])
	binary.BigEndian.PutUint64(s[24:32], x.s[3])
	binary.BigEndian.PutUint64(s[32:40], x.s[4])
	binary.BigEndian.PutUint64(s[40:48], x.s[5])
	binary.BigEndian.PutUint64(s[48:56], x.s[6])
	binary.BigEndian.PutUint64(s[56:64], x.s[7])
	return s
}

func (x *Xoshiro512p) SetState(state []byte) {
	mix := NewSplitMix64(time.Now().UTC().UnixNano())
	x.s[0] = bytesToState64(state, 0, &mix)
	x.s[1] = bytesToState64(state, 1, &mix)
	x.s[2] = bytesToState64(state, 2, &mix)
	x.s[3] = bytesToState64(state, 3, &mix)
	x.s[4] = bytesToState64(state, 4, &mix)
	x.s[5] = bytesToState64(state, 5, &mix)
	x.s[6] = bytesToState64(state, 6, &mix)
	x.s[7] = bytesToState64(state, 7, &mix)
}

func (x *Xoshiro512p) Seed(seed int64) {
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

func (x *Xoshiro512p) Uint64() uint64 {
	result := x.s[0] + x.s[2]
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

func (x *Xoshiro512p) Int64() int64 {
	return unsafeUint64ToInt64(x.Uint64())
}

func (x *Xoshiro512p) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

func (x *Xoshiro512p) Float64() float64 {
	return math.Float64frombits(0x3ff<<52|x.Uint64()>>12) - 1.0
}

// Call Uint64() * 2^256
func (x *Xoshiro512p) Jump() {
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

	x.s[0] = s[0]
	x.s[1] = s[1]
	x.s[2] = s[2]
	x.s[3] = s[3]
	x.s[4] = s[4]
	x.s[5] = s[5]
	x.s[6] = s[6]
	x.s[7] = s[7]
}

// Call Uint64() * 2^192
func (x *Xoshiro512p) LongJump() {
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

	x.s[0] = s[0]
	x.s[1] = s[1]
	x.s[2] = s[2]
	x.s[3] = s[3]
	x.s[4] = s[4]
	x.s[5] = s[5]
	x.s[6] = s[6]
	x.s[7] = s[7]
}

func (x Xoshiro512p) String() string {
	return fmt.Sprintf("%0128x", x.State())
}

func (x Xoshiro512p) GoString() string {
	return "xrand.Xoshiro512p{state:\"" + x.String() + "\"}"
}
