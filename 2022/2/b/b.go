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

	pts := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line[0] == 'A' {
			if line[2] == 'X' {
				pts += 3
			} else if line[2] == 'Y' {
				pts += 4
			} else {
				pts += 8
			}
		} else if line[0] == 'B' {
			if line[2] == 'X' {
				pts += 1
			} else if line[2] == 'Y' {
				pts += 5
			} else {
				pts += 9
			}
		} else {
			if line[2] == 'X' {
				pts += 2
			} else if line[2] == 'Y' {
				pts += 6
			} else {
				pts += 7
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pts)
}
