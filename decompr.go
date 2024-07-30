package main

import (
	"bufio"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"os"
	"strconv"
)

func Decompress() {
	file, _ := os.Open("test.huff")

	huffmanMap := make(map[string]string)

	reader := bufio.NewReader(file)

	bits := bitsToConsider(file)

	fmt.Println("Bits to consider:", bits)

	line, _, _ := reader.ReadLine()
	Must(msgpack.Unmarshal(line, &huffmanMap))

	createFile, _ := os.Create("test.huff.decomp")
	writer := bufio.NewWriter(createFile)
	keyBuffer := ""

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanBytes)

	counter := uint(0)

label:
	for scanner.Scan() {
		b := scanner.Text()

		for i := 0; i < 8; i++ {
			if counter == bits {
				break label
			}

			bit := (b[0] >> uint(7-i)) & 1
			keyBuffer += string(bit + '0')

			if char, ok := huffmanMap[keyBuffer]; ok {
				if char == "\\n" {
					char = "\n"
				}
				Must2(writer.WriteString(char))

				keyBuffer = ""
			}

			counter++
		}
	}

	fmt.Println()

	fmt.Println("Counter:", counter)

	Must(writer.Flush())
}

func bitsToConsider(file *os.File) uint {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()

	// read until reaching \n
	numStr := ""
	i := 1
	for {
		buff := make([]byte, 1)
		Must2(file.ReadAt(buff, fileSize-int64(i)))

		if string(buff) == "\n" {
			break
		}

		i++
		numStr = string(buff) + numStr
	}

	atoi, _ := strconv.Atoi(numStr)

	return uint(atoi)
}
