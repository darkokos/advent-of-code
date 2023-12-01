package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var matrix [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var s []int
		for i := 0; i < len(line); i++ {
			s = append(s, int(line[i]))
		}

		matrix = append(matrix, s)
	}

	num := 0
	for i, s := range matrix {
		for j, e := range s {
			if i == 0 || j == 0 || i == len(matrix)-1 || j == len(s)-1 {
				num++
				continue
			}

			breaker := 0
			for k := 0; k < j; k++ {
				if matrix[i][k] >= e {
					breaker++
					break
				}
			}

			for k := j + 1; k < len(s); k++ {
				if matrix[i][k] >= e {
					breaker++
					break
				}
			}

			for k := 0; k < i; k++ {
				if matrix[k][j] >= e {
					breaker++
					break
				}
			}

			for k := i + 1; k < len(matrix); k++ {
				if matrix[k][j] >= e {
					breaker++
					break
				}
			}

			if breaker == 4 {
				continue
			}

			num++
		}
	}

	fmt.Println(num)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
