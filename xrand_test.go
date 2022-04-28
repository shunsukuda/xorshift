package xrand

import "testing"

func BenchmarkSplitMix64Seed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	var x SplitMix64
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}
func BenchmarkSplitMix64Uint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewSplitMix64(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkSplitMix64Int64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewSplitMix64(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkSplitMix64Int63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewSplitMix64(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro128pSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoroshiro128pUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoroshiro128pInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoroshiro128pInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro128pFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoroshiro128pJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoroshiro128pLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoroshiro128ppSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoroshiro128ppUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoroshiro128ppInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoroshiro128ppInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro128ppFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoroshiro128ppJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoroshiro128ppLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoroshiro128ssSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoroshiro128ssUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoroshiro128ssInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoroshiro128ssInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro128ssFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoroshiro128ssJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoroshiro128ssLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro128ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoshiro256pSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoshiro256pUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoshiro256pInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoshiro256pInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoshiro256pFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoshiro256pJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoshiro256pLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoshiro256ppSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoshiro256ppUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoshiro256ppInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoshiro256ppInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoshiro256ppFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoshiro256ppJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoshiro256ppLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoshiro256ssSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoshiro256ssUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoshiro256ssInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoshiro256ssInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoshiro256ssFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoshiro256ssJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoshiro256ssLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro256ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoshiro512pSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoshiro512pUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoshiro512pInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoshiro512pInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoshiro512pFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoshiro512pJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoshiro512pLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512p(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoshiro512ppSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoshiro512ppUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoshiro512ppInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoshiro512ppInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoshiro512ppFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoshiro512ppJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoshiro512ppLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x := NewXoshiro512pp(int64(1234567890))
		x.LongJump()
	}
}

func BenchmarkXoshiro512ssSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoshiro512ssUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoshiro512ssInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoshiro512ssInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoshiro512ssFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoshiro512ssJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoshiro512ssLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoshiro512ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoroshiro1024ppSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoroshiro1024ppUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoroshiro1024ppInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoroshiro1024ppInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro1024ppFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoroshiro1024ppJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoroshiro1024ppLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024pp(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoroshiro1024sSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoroshiro1024sUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoroshiro1024sInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoroshiro1024sInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro1024sFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoroshiro1024sJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoroshiro1024sLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024s(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}

func BenchmarkXoroshiro1024ssSeed(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Seed(int64(i))
	}
}

func BenchmarkXoroshiro1024ssUint64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Uint64()
	}
}

func BenchmarkXoroshiro1024ssInt64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int64()
	}
}

func BenchmarkXoroshiro1024ssInt63(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Int63()
	}
}

func BenchmarkXoroshiro1024ssFloat64(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		_ = x.Float64()
	}
}

func BenchmarkXoroshiro1024ssJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.Jump()
	}
}
func BenchmarkXoroshiro1024ssLongJump(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	x := NewXoroshiro1024ss(int64(1234567890))
	for i := 0; i < b.N; i++ {
		x.LongJump()
	}
}
