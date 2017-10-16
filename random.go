package main

import (
	"bufio"
	"os"
	"math/rand"
	"fmt"
	"strconv"
)

func main() {
	outputFile, outputError := os.OpenFile("millionRandom.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("File open failed.")
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	for i := 0; i < 100000; i++ {
		str := strconv.Itoa(rand.Intn(1000))
		outputWriter.WriteString(str + "\n")
	}

}
