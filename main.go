package main

import "fmt"
import "github.com/libercv/peerbackup/crawler"

func main() {
	fmt.Println("Peer backup")
	config := ReadConfig()
	fmt.Printf("Source dir: %s\n", config.SrcDir)
	fmt.Printf("Destination dir: %s\n", config.DstDir)
	crawler.SyncDir(config.SrcDir, config.DstDir, false)
}
