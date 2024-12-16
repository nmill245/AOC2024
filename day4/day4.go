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

func parseXmasLines(lines []string) int {
	num := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if line[j] == 'X' {
				//check backwards
				if j > 2 && line[j-1] == 'M' && line[j-2] == 'A' && line[j-3] == 'S' {
					num++
				}
				//check forwards
				if j < len(line)-3 && line[j+1] == 'M' && line[j+2] == 'A' && line[j+3] == 'S' {
					num++
				}
				//check down
				if i < len(lines)-3 && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					num++
				}
				//check up
				if i > 2 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					num++
				}
				//check up-right
				if i > 2 && j < len(line)-3 && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					num++
				}
				//check down left
				if i < len(lines)-3 && j > 2 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					num++
				}
				//check down right
				if i < len(lines)-3 && j < len(lines)-3 && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					num++
				}
				//check up left
				if i > 2 && j > 2 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					num++
				}
			}
		}
	}
	return num
}

func parseMasLines(lines []string) int {
	num := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if line[j] == 'M' {
				if i > 1 && j > 1 && lines[i-1][j-1] == 'A' && lines[i-2][j-2] == 'S' {
					if (line[j-2] == 'S' && lines[i-2][j] == 'M') || (line[j-2] == 'M' && lines[i-2][j] == 'S') {
						num++
					}
				}
				if i < len(lines)-2 && j < len(lines)-2 && lines[i+1][j+1] == 'A' && lines[i+2][j+2] == 'S' {
					if (line[j+2] == 'M' && lines[i+2][j] == 'S') || (line[j+2] == 'S' && lines[i+2][j] == 'M') {
						num++
					}
				}
			}

		}
	}

	return num
}

func main() {
	data := readFile("input.txt")
	data = strings.Trim(data, "\r\n")
	lines := strings.Split(data, "\r\n")
	result := parseMasLines(lines)
	fmt.Println(result)
}
