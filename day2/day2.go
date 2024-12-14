package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseLine(line string) []int {
	final := []int{}
	data := strings.Split(line, " ")
	for i := 0; i < len(data); i++ {
		num, err := strconv.Atoi(data[i])
		if err != nil {
			panic(err)
		}
		final = append(final, num)
	}
	return final
}

func main() {
	data := readFile("test_input.txt")
	data = strings.Trim(data, "\r\n")
	lines := strings.Split(data, "\r\n")
	for i := 0; i < len(lines); i++ {
		values := parseLine(lines[i])
		fmt.Println(values)
	}

}
