package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type config struct {
	SrcDir string
	DesDir string
	Files  []string
}

func convertDzhSectorToWFT(srcFile, desDir string) {
	//Open the DZH sector file
	file, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//generate des file name
	_, fileName := filepath.Split(srcFile)
	fileExt := filepath.Ext(fileName)
	fileNameNew := strings.TrimSuffix(fileName, fileExt) + ".txt"

	//create des file
	desFile, err := os.Create(filepath.Join(desDir, fileNameNew))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer desFile.Close()

	//create a reader buffer & skip file head
	reader := bufio.NewReader(file)
	reader.Discard(4)

	//make a buffer for stock code, every stock length is 12 byte
	buf := make([]byte, 12)

	//read stock code one by one
	for {
		n, _ := reader.Read(buf)
		if n != 12 {
			break
		}
		s := string(buf[0:8])
		windName := s[2:] + "." + s[0:2]

		num, err := desFile.WriteString(windName + "\n")
		if err != nil {
			fmt.Println(err, num)
		}
	}

	desFile.Sync()
}

func main() {
	stream, err := ioutil.ReadFile("convertDzhSectorToWFT.json")
	if err != nil {
		log.Fatal("can't open config file")
	}

	cfg := &config{}
	json.Unmarshal(stream, cfg)

	for _, n := range cfg.Files {
		convertDzhSectorToWFT(filepath.Join(cfg.SrcDir, n), cfg.DesDir)
	}
}
