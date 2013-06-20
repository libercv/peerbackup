package main

import (
	"os"
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

