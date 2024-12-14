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

func merge(left []int, right []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		final = append(final, left[i])
	}
	for ; j < len(right); j++ {
		final = append(final, right[j])
	}
	return final
}

func sort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	left := sort(data[:len(data)/2])
	right := sort(data[len(data)/2:])
	return merge(left, right)
}

func difference(larr []int, rarr []int) int {
	result := 0
	for i := 0; i < len(larr); i++ {
		diff := larr[i] - rarr[i]
		if diff < 0 {
			diff = diff * -1
		}
		result = result + diff
	}
	return result
}

func hash(data []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(data); i++ {
		m[data[i]] = m[data[i]] + 1
	}
	return m
}

func sim(larr []int, count map[int]int) int {
	result := 0
	for i := 0; i < len(larr); i++ {
		result = result + larr[i]*count[larr[i]]
	}
	return result
}

func main() {
	data := readFile("input.txt")
	trimmedData := strings.Trim(string(data), "\r\n")
	splitData := strings.Split(trimmedData, "\r\n")
	larr, rarr := parseData(splitData)
	countarr := hash(rarr)
	result := sim(larr, countarr)
	fmt.Printf("Final similarity %d\n", result)

	/**
	*larr = sort(larr)
	*rarr = sort(rarr)
	diff := difference(larr, rarr)
	fmt.Printf("Final difference: %d\n", diff)
	**/

}
