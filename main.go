package main

import (
	//"path"
	//"bufio"
	"fmt"
	//"io"
	"os"
	"github.com/libercv/peerbackup/hasher"
	"github.com/libercv/peerbackup/crawler"
)

func main() {
	fmt.Println("Peer backup")
	config := ReadConfig()
	fmt.Printf("Source dir: %s\n", config.SrcDir)
	fmt.Printf("Destination dir: %s\n", config.DstDir)

	fi, err := os.Open("peerbackup")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	/*
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
			fmt.Printf("%x %x\n", hasher.GetSHA256(buf), hasher.GetAdler32(buf))
			WriteFileGZIP(path.Join(config.DstDir, fmt.Sprintf("%x",hasher.GetSHA256(buf))), buf)
		}
	*/
	ch := crawler.WalkDir(config.SrcDir)
	for src := range ch {
		fmt.Printf(src.Name)
		fi, _:=os.Stat(src.Name)
		if !fi.IsDir() {
			fmt.Printf("  %x\n", hasher.GetFileSHA256(src.Name))
		}
	}

}
