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

	num, max := 0, 0
	for i, s := range matrix {
		for j, e := range s {
			breaker := 0
			ss1, ss2, ss3, ss4, ss := j, len(s)-j-1, i, len(matrix)-i-1, 0
			for k := j - 1; k >= 0; k-- {
				if matrix[i][k] >= e {
					ss1 = j - k
					breaker++
					break
				}
			}

			for k := j + 1; k < len(s); k++ {
				if matrix[i][k] >= e {
					ss2 = k - j
					breaker++
					break
				}
			}

			for k := i - 1; k >= 0; k-- {
				if matrix[k][j] >= e {
					ss3 = i - k
					breaker++
					break
				}
			}

			for k := i + 1; k < len(matrix); k++ {
				if matrix[k][j] >= e {
					ss4 = k - i
					breaker++
					break
				}
			}

			if breaker == 4 {
				continue
			}

			ss = ss1 * ss2 * ss3 * ss4
			fmt.Println(ss1, ss2, ss3, ss4)
			if ss > max {
				max = ss
			}
			num++
		}
	}

	fmt.Println(num)
	fmt.Println(max)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
