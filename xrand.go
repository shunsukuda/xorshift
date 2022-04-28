package xrand

import (
	"encoding/binary"
	"unsafe"
)

func bytesToState64(b []byte, pos int, mix *SplitMix64) uint64 {
	var tmp []byte
	b = b[:cap(b)]
	l := len(b)
	pos += 1
	if pos <= 0 || l <= 8*(pos-1) {
		return mix.Uint64()
	}
	if len(b) >= 8*pos {
		tmp = b[8*(pos-1) : 8*pos]
	} else {
		tmp = make([]byte, 8)
		copy(tmp, b[8*(pos-1):l])
	}
	return binary.BigEndian.Uint64(tmp)
}

func unsafeInt64ToUint64(v int64) uint64 {
	return *(*uint64)(unsafe.Pointer(&v))
}

func unsafeUint64ToInt64(v uint64) int64 {
	return *(*int64)(unsafe.Pointer(&v))
}
