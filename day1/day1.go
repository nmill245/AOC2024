package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func parseData(data []string) ([]int, []int) {
	length := len(data)
	larr := make([]int, length)
	rarr := make([]int, length)
	var err error
	for i := 0; i < length; i++ {
		line := strings.Split(data[i], "   ")
		larr[i], err = strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		rarr[i], err = strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}
	}
	return larr, rarr
}
func main() {
	data := readFile("test_input.txt")
	trimmedData := strings.Trim(string(data), "\n")
	splitData := strings.Split(trimmedData, "\n")
	larr, rarr := parseData(splitData)
	fmt.Println(larr)
	fmt.Println(rarr)

}
