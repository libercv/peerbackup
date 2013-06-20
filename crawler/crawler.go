//  This file is part of Peerbackup
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

package crawler

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
)

// entry contains information about the file
// or directory
type Entry struct {
	Name string
	info os.FileInfo
}

// We create channels of communication between coroutines
var (
	chSrc = make(chan Entry)
	chDst = make(chan Entry)
	chOutput chan Entry
)

func WalkDir(rootSrc string)  (chan Entry) {
	chOutput =  make (chan Entry)
	go walkDirAsync(rootSrc)
	return chOutput
}

func walkDirAsync(rootSrc string) {
	filepath.Walk(rootSrc, visitFile)
	close(chOutput)
}

func visitFile(path string, f os.FileInfo, err error) error {
	// TODO: Handle errors
	chOutput <- Entry{path, f}
	return nil
}

// SyncDir synchronizes two directories
func SyncDir(rootSrc string, rootDst string, rm_deleted bool) {
	go walkSrc(rootSrc)
	go walkDst(rootDst)
	lenSrc := len(rootSrc)
	lenDst := len(rootDst)
	dst, ok := <-chDst
	for src := range chSrc {
		// Name of the destination file to be copied/created
		fileDst := filepath.Join(rootDst, src.Name[lenSrc:])
		if ok {
		// We still have destination files to check
			if rm_deleted {
			// Check for files to be deleted
				for src.Name[lenSrc:] > dst.Name[lenDst:] {
					os.Remove(dst.Name)
					fmt.Printf("Erased %s \n", dst.Name)
					dst, ok = <-chDst
				}
			}
			if  src.Name[lenSrc:] == dst.Name[lenDst:] {
				// File exists in both dirs. Copy only if
				// newer version exists
				if src.info.ModTime().After(dst.info.ModTime()) {
					entryCopy(fileDst, src.Name, src.info)
				}
				dst, ok = <-chDst
			} else if src.Name[lenSrc:] < dst.Name[lenDst:] {
				// Copy a new file
				entryCopy(fileDst, src.Name, src.info)
			}
		} else  { // !ok
			// No more files in the destination directory. 
			// Copy a new file
			entryCopy(fileDst, src.Name, src.info)
		}
	}
}

func walkSrc(pathSrc string) {
	// TODO: Handle errors
	filepath.Walk(pathSrc, visitSrc)
	close(chSrc)
}

func walkDst(path string) {
	// TODO: Handle errors
	filepath.Walk(path, visitDst)
	close(chDst)
}

func visitSrc(path string, f os.FileInfo, err error) error {
	// TODO: Handle errors
	chSrc <- Entry{path, f}
	return nil
}

func visitDst(path string, f os.FileInfo, err error) error {
	// TODO: Handle erros
	chDst <- Entry{path, f}
	return nil
}

// entryCopy copies an entry, whether it's a file or a directory
func entryCopy(dstName, srcName string, fi os.FileInfo) (written int64, err error) {
	// TODO: Handle errors
	written = 0
	err = nil

	if fi.Mode().IsDir() {
		// Create a directory
		os.Mkdir(dstName, fi.Mode())
		fmt.Printf("Directory %s created\n", dstName)
	} else {
		// Copy a file
		written, err = fileCopy(dstName, srcName, fi.Mode())
		fmt.Printf("Copied %s\n", dstName)
	}

	os.Chtimes(dstName, fi.ModTime(), fi.ModTime())
	return written, err
}

// fileCopy copies a file, maintaining file attributes (FileMode)
func fileCopy(dstName, srcName string, srcMode os.FileMode) (written int64, err error) {
	// TODO: Handle errors

	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	defer os.Chmod(dst.Name(), srcMode)

	return io.Copy(dst, src)
}
