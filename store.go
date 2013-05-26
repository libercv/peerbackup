package main

import (
	"os"
	"code.google.com/p/lzma"
	"compress/gzip"
)

func WriteFileLZMA(filename string, buf []byte) {
	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	w := lzma.NewWriterLevel(fi, lzma.BestCompression)
	w.Write(buf)
	w.Close()
}

func WriteFileGZIP(filename string, buf []byte) {
	fi, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	w := gzip.NewWriter(fi)
	w.Write(buf)
	w.Close()
}