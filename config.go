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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	settingsFile = "settings.json"
)

// Config contains configuration of the program
type Config struct {
	SrcDir string // Source directory of the backup
	DstDir string // Destination directory of the backup
}

func WriteJSONConfig(conf *Config) {
	file, err := os.Create(settingsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cadena, _ := json.MarshalIndent(conf, "", "  ")
	file.Write(cadena)
	file.WriteString("\n")
}

func ReadJSONConfig() Config {
	var conf Config
	archivo, _ := ioutil.ReadFile(settingsFile)
	err := json.Unmarshal(archivo, &conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return conf
}
