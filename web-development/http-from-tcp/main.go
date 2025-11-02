package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	reader, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", inputFilePath, err.Error())
	}

	for line := range getLinesChannel(reader) {
		println("read:", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)
		line := ""
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if line != "" {
					ch <- line
				}
				if errors.Is(err, io.EOF) {
					return
				}
				ch <- fmt.Sprintf("Error reading file: %s\n", err.Error())
				return
			}

			str := string(b[:n])
			lineParts := strings.Split(str, "\n")
			for i := 0; i < len(lineParts)-1; i++ {
				ch <- line + lineParts[i]
				line = ""
			}
			line += lineParts[len(lineParts)-1]
		}
	}()

	return ch
}
