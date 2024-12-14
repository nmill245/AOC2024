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

func checkSafety(line []int) bool {
	num1 := line[0]
	num2 := line[1]
	ascending := false
	if num1 < num2 {
		ascending = true
		if num2-num1 > 3 {
			return false
		}
	} else if num1 > num2 {
		ascending = false
		if num1-num2 > 3 {
			return false
		}
	} else {
		return false
	}
	for i := 2; i < len(line); i++ {
		num1 = line[i-1]
		num2 = line[i]
		if num1 < num2 {
			if !ascending {
				return false
			} else if num2-num1 > 3 {
				return false
			}
		} else if num1 > num2 {
			if ascending {
				return false
			}
			if num1-num2 > 3 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	data := readFile("input.txt")
	data = strings.Trim(data, "\r\n")
	lines := strings.Split(data, "\r\n")
	found := false
	num := 0
	for i := 0; i < len(lines); i++ {
		values := parseLine(lines[i])
		safe := checkSafety(values)
		if safe {
			num++
			found = true
		}
		if !found {
			for j := 0; j < len(values); j++ {
				new_values := []int{}
				new_values = append(new_values, values[:j]...)
				new_values = append(new_values, values[j+1:]...)
				safe = checkSafety(new_values)
				if safe {
					num++
					break
				}
				fmt.Println(new_values)
			}
		}
		found = false
	}
	fmt.Println("Number of safe:", num)
}
