package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func divideString(data string) []string {
	s := strings.ReplaceAll(data, ")", ")\n")
	k := strings.Split(s, "\n")
	return k

}
func checkdo(line string) bool {
	match, err := regexp.MatchString(`do\(\)`, line)
	if err != nil {
		panic(err)
	}
	return match
}
func checkdont(line string) bool {
	rgx := regexp.MustCompile(`don\'t\(\)`)
	rs := rgx.FindStringSubmatch(line)
	if len(rs) > 0 {
		return true
	}
	return false
}

func parseLine(line string) int {
	rgx := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	rs := rgx.FindStringSubmatch(line)
	if len(rs) < 3 {
		return 0
	}
	num1, er1 := strconv.Atoi(rs[1])
	if er1 != nil {
		panic(er1)
	}
	num2, er2 := strconv.Atoi(rs[2])
	if er2 != nil {
		panic(er2)
	}
	return num1 * num2
}

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	data := readFile("input.txt")
	data = strings.Trim(data, "\r\n")
	lines := divideString(data)
	result := 0
	allowed := true
	for i := 0; i < len(lines); i++ {
		if allowed {
			allowed = !checkdont(lines[i])
			if allowed {
				value := parseLine(lines[i])
				result += value
			}
		} else {
			allowed = checkdo(lines[i])
			if allowed {
				value := parseLine(lines[i])
				result += value
			}
		}
	}
	fmt.Println(result)
}
