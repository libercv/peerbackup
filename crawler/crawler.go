package crawler

import (
	"fmt"
	"os"
	"path/filepath"
)

type entry struct {
	name string
	info *os.FileInfo
}

var (
	chSrc = make(chan entry)
	chDst = make(chan entry)
)

func SyncDir(rootSrc string, rootDst string, rm_deleted bool) {
	go walkSrc(rootSrc)
	go walkDst(rootDst)
	lenSrc := len(rootSrc)
	lenDst := len(rootDst)
	dst, ok := <-chDst
	for src := range chSrc {
		for src.name[lenSrc:] > dst.name[lenDst:] {
			fmt.Printf("Erase %s \n", dst.name)
			dst, ok = <-chDst
		}
		if ok && src.name[lenSrc:] == dst.name[lenDst:] {
			fmt.Printf("file ok: %s\n", src.name[lenSrc:])
			dst, ok = <-chDst
		} else if !ok {
			fmt.Printf("Create %s\n", src.name)
		} else if src.name[lenSrc:] < dst.name[lenDst:] {
			fmt.Printf("Create %s, waiting %s\n", src.name, dst.name)
		}
	}
}

func walkSrc(pathSrc string) {
	filepath.Walk(pathSrc, visitSrc)
	close(chSrc)
}

func walkDst(path string) {
	filepath.Walk(path, visitDst)
	close(chDst)
}

func visitSrc(path string, f os.FileInfo, err error) error {
	chSrc <- entry{path, &f}
	return nil
}

func visitDst(path string, f os.FileInfo, err error) error {
	chDst <- entry{path, &f}
	return nil
}
