package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

	//create a reader buffer
	reader := bufio.NewReader(file)
	//skip file head
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
	convertDzhSectorToWFT("c:\\DZH2\\USERDATA\\block\\A多头.blk", "r:\\")
	convertDzhSectorToWFT("c:\\DZH2\\USERDATA\\block\\A间距.blk", "r:\\")
	convertDzhSectorToWFT("c:\\DZH2\\USERDATA\\block\\A剑来.blk", "r:\\")
	convertDzhSectorToWFT("c:\\DZH2\\USERDATA\\block\\A周线放量.blk", "r:\\")
}
