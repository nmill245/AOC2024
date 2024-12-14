package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseLine(line []string) []int {
	final := []int{}
	data := strings.Split(line, " ")
	for i := 0; i < len(data); i++ {
		final = append(final, AtoI(data[i]))
	}
	return final
}

func main() {
	data := readFile("test_input.txt")
	data = strings.Trim(data, "\r\n")
	lines := strings.Split(data, "\r\n")
	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}
