package main

import (
	"fmt"
	"os"
)

func main() {
	var fileName string = "config.txt"
	readFile := func(fileName string) ([]byte, int, error) {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err.Error())
			return nil, 0, err
		}
		defer file.Close()
		bytesToRead := 256
		data := make([]byte, bytesToRead)
		numOfReadedBytes, err := file.Read(data)
		return data, numOfReadedBytes, err
	}
	for i := 0; i < 20; i++ {
		data, readed, err := readFile(fileName)
		if err != nil {
			break
		}
		if readed > 0 {
			fmt.Println(string(data[:readed]))
		}
	}
}
