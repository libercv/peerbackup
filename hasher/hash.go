package hasher

import (
	"crypto/sha256"
//	"hash/adler32"
)

func GetSHA256(buf []byte) ([]byte) {
	h := sha256.New()
	h.Write(buf)
	return h.Sum(nil)
}

func GetAdler32(buf []byte) (uint32) {
	ad:= NewRollingAdler32()
	ad.Write(buf)
	return ad.Sum32()
}

