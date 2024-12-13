package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createArray(num int) []int {
	slice := make([]int, num)
	for i := 0; i < num; i++ {
		slice[i] = i
	}
	return slice
}

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	message := fmt.Sprintf("Hello World")
	fmt.Println(message)
	slice := createArray(6)
	fmt.Println(slice)
	data := readFile("test_input.txt")
	splitData := strings.Split(string(data), "\n")
	fmt.Printf("File content: \n")
	for i := 0; i < len(splitData); i++ {
		fmt.Printf("%s\n", splitData[i])
	}
	bytetoInt, err := strconv.Atoi(string(data[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println(bytetoInt)
}
