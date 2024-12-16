package main

import (
	"fmt"
	"os"
	"sort"
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

func parseRules(rules []string) map[int][]int {
	rulehash := make(map[int][]int)
	for _, rule := range rules {
		vals := strings.Split(rule, "|")
		num1, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		num2, er2 := strconv.Atoi(vals[1])
		if er2 != nil {
			panic(er2)
		}
		rulehash[num1] = append(rulehash[num1], num2)

	}
	for _, num := range rulehash {
		sort.Ints(num)
	}
	return rulehash
}

func contains(value int, list []int) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := right - (right-left)/2
		if list[mid] == value {
			return mid
		} else if list[mid] > value {
			right = mid - 1
		} else if list[mid] < value {
			left = mid + 1
		}
	}
	return -1

}

func parseLines(lines []string, rules map[int][]int) int {
	result := 0
	for _, line := range lines {
		vals := strings.Split(line, ",")
		seen := []int{}
		bad := false
		found := false
		for _, val := range vals {
			bad = false
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			for i, old := range seen {
				if contains(old, rules[num]) != -1 {
					for j := i; j < len(seen); j++ {
						temp := seen[j]
						seen[j] = num
						num = temp
					}
					seen = append(seen, num)
					bad = true
					found = true
					break
				}
			}
			if !bad {
				seen = append(seen, num)
			}
		}
		if found {
			fmt.Println(seen)
			mid := len(seen) / 2
			result += seen[mid]
		}
	}
	return result
}

func main() {
	//file := "test_input.txt"
	file := "input.txt"
	data := readFile(file)
	data = strings.Trim(data, "\r\n")
	lines := strings.Split(data, "\r\n")
	rules := []string{}
	updates := []string{}
	for _, line := range lines {
		if strings.Index(line, "|") != -1 {
			rules = append(rules, line)

		} else if strings.Index(line, ",") != -1 {
			updates = append(updates, line)
		}
	}
	hashrules := parseRules(rules)
	result := parseLines(updates, hashrules)
	fmt.Println("Final Result: ", result)
}
