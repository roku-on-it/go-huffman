package main

import (
	"bufio"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"os"
	"shrink/huffman"
)

func Compress() {
	freqMap := make(map[string]uint64)

	file, _ := os.Open("test")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		char := scanner.Text()
		freqMap[char]++
	}

	Must(file.Close())

	root := huffman.BuildHuffmanTree(freqMap)
	huffmanMap := huffman.BuildHuffmanMap(root)

	file, _ = os.Open("test")
	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	file, _ = os.Create("test.huff")
	writer := bufio.NewWriter(file)

	mapToWrite := make(map[string]string)

	for k, v := range huffmanMap {

		if k == "\n" {
			k = "\\n"
		}

		mapToWrite[v] = k
	}

	msgPackBytes, err := msgpack.Marshal(mapToWrite)
	if err != nil {
		panic(err)
	}

	Must2(file.Write(msgPackBytes))
	Must2(writer.WriteString("\n"))

	char := 0
	counter := 0
	var bitsBuffer []int
	bitsToConsider := 0

	for scanner.Scan() {
		bits := huffmanMap[scanner.Text()]

		for _, bit := range bits {
			bitInt := int(bit - '0')

			char = char | bitInt
			counter++
			bitsBuffer = append(bitsBuffer, bitInt)
			bitsToConsider++

			if counter == 8 {
				Must(writer.WriteByte(byte(char)))

				char = 0
				counter = 0
				bitsBuffer = bitsBuffer[:0]
			} else {
				char = char << 1
			}

		}
	}

	if counter != 0 {
		remaining := 8 - counter
		bits := append(bitsBuffer, make([]int, remaining)...)

		var char byte
		for i, bit := range bits {
			char = char | byte(bit)
			if i != 7 {
				char = char << 1
			}
		}

		Must(writer.WriteByte(char))
	}

	Must(writer.Flush())

	Must2(file.WriteString(fmt.Sprintf("\n%d", bitsToConsider)))

	Must(file.Close())
}
