//	This file is part of Peerbackup
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU General Public License as published by
//  the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU General Public License for more details.
//
//  You should have received a copy of the GNU General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.

// TODO: use static variables for the config file name and token names.

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

// config contains configuration of the program
type config struct {
	srcDir, dstDir string
}

// ReadConfig reads the configuration file and returns it in the form
// of a "config" struct
func ReadConfig() config {
	var conf config
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if isPrefix {
			log.Fatal("Error: Unexpected long line reading", file.Name())
		}
		parseConfigLine(string(line), &conf)
	}
	return conf
}

// parseConfigLine parses a string containing a line of the configuration
// file extracting its information and filling a variable of the "config" struct
func parseConfigLine(line string, conf *config) {
	tokens := strings.Split(line, "=")
	switch tokens[0] {
	case "src":
		conf.srcDir = tokens[1]
	case "dst":
		conf.dstDir = tokens[1]
	}
}
