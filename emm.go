package main

import (
	"fmt"
	//"log"
	"os"
)

func openFile() *os.File {
	//log.Print("FileOpen Started.")
	//defer log.Print("FileOpen Ended.")

	fmt.Print("Type in the path of the input file:\n")

	var fileInput string
	fmt.Scanf("%s",&fileInput)

	file,err := os.OpenFile(fileInput, os.O_RDONLY, 0666)
	if err != nil {
		panic("file open failure")
	}

	return file
}

func scanFile(file *os.File, m *map[uint32]uint32) {
	//log.Print("File Scan Started.")
	//defer log.Print("File Scan Ended.")

	var numScan uint32
	for {
		n, err := fmt.Fscanf(file,"%d\n",&numScan)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				panic("read error")
			}
		}
		if n==1 {
			(*m)[numScan]++
		}
	}
}

func renderMap(m *map[uint32]uint32) {
	//log.Print("Render started.")
	//defer log.Print("Render Ended.")

	for k,v := range(*m) {
		fmt.Printf("%d occured %d time(s).\n",k,v)
	}
}

func main() {
	file := openFile()
	defer file.Close()

	mNum := make(map[uint32]uint32)
	scanFile(file, &mNum)

	renderMap(&mNum)
}
