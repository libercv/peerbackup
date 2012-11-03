package main

import "fmt"

func main() {
	fmt.Println("Peer backup")
	config := ReadConfig()
	fmt.Printf("Source dir: %s\n", config.SrcDir)
	fmt.Printf("Destination dir: %s\n", config.DstDir)
}
