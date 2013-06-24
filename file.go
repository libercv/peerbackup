package main

import (
	"os"
	"fmt"
	"bufio"
	"path"
	"io"
	"container/list"
	"github.com/libercv/peerbackup/hasher"
	"path/filepath"
)

type FileMetadata struct {
	FileInfo os.FileInfo
	Path string
	CryptHash []byte
	AdlerHash uint32
	// List of hash fragments 
	Fragments list.List
}


type FileBlock struct {
	CryptHash []byte
	AdlerHash uint32
}

func GetFileInfo(fileName string) *FileMetadata {
	m := new(FileMetadata)

	fi, _:=os.Stat(fileName)
	m.FileInfo=fi

	m.Path=filepath.Dir(fileName)

	// Inefficient way of calculating the hashes
	// Multiwriter should be better. 
	if !fi.IsDir() {
		m.CryptHash=hasher.GetFileSHA256(fileName)
		m.AdlerHash=hasher.GetFileAdler32(fileName)
	}

	// Again, use multiwriter, not parse thrice every file...
	block:=new(FileBlock)
	m.Fragments.PushBack(block)
	return m
}


func (fm *FileMetadata) BackupFile(dstFolder string) {

	name := path.Join(fm.Path, fm.FileInfo.Name())

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
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		hash:=hasher.GetSHA256(buf)
		WriteFileGZIP(path.Join(dstFolder, fmt.Sprintf("%x", hash)), buf)
	}
}
