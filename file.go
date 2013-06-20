package main

import "os"

type FileMetadata struct {
	os.FileInfo
	CryptHash []byte
	AdlerHash uint32
	// List of hash fragments 
}
