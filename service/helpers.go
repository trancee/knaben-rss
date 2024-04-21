package service

import (
	"unsafe"

	"golang.org/x/crypto/blake2b"
)

func b2s(b []byte) string {
	// Ignore if your IDE shows an error here; it's a false positive.
	p := unsafe.SliceData(b)
	return unsafe.String(p, len(b))
}

// https://josestg.medium.com/140x-faster-string-to-byte-and-byte-to-string-conversions-with-zero-allocation-in-go-200b4d7105fc
func s2b(s string) []byte {
	p := unsafe.StringData(s)
	b := unsafe.Slice(p, len(s))
	return b
}

// func shortDuration(d time.Duration) string {
// 	s := d.String()

// 	if strings.HasSuffix(s, "m0s") {
// 		s = s[:len(s)-2]
// 	}

// 	if strings.HasSuffix(s, "h0m") {
// 		s = s[:len(s)-2]
// 	}

// 	return s
// }

//	func ShortDuration(s string) string {
//		if d, err := time.ParseDuration(s + "s"); err != nil {
//			panic(err)
//		} else {
//			return shortDuration(d)
//		}
//	}

// func ShortDuration(s int) string {
// 	d := time.Duration(s) * time.Second

// 	return shortDuration(d)
// }

func Hash(data []byte) []byte {
	hash := blake2b.Sum256(data)
	return hash[:]
}
