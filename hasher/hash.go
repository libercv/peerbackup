package hasher

import (
	"crypto/sha256"
	"os"
	"bufio"
	"io"
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

func GetFileSHA256(name string) ([]byte) {
	fi, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	r := bufio.NewReader(fi)
	buf := make([]byte, 262144)
	h:=sha256.New()
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		h.Write(buf)
	}
	return h.Sum(nil)
}


