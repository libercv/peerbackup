package main

import (
	//"path"
	//"bufio"
	"fmt"
	//"io"
	//"os"
	//"github.com/libercv/peerbackup/hasher"
	"github.com/libercv/peerbackup/crawler"
)

func main() {
	fmt.Println("Peer backup")
	config := ReadConfig()
	ch := crawler.WalkDir(config.SrcDir)
	for src := range ch {
		f := GetFileInfo(src.Name)
		fmt.Printf("%s - %s", f.Path, f.FileInfo.Name())
		if !f.FileInfo.IsDir() {
			fmt.Printf("  %x   %d", f.CryptHash, f.AdlerHash)
			f.BackupFile(config.DstDir)
		}
		fmt.Printf("\n")
	}

}
