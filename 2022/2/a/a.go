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

		if line[2] == 'X' {
			pts += 1
			if line[0] == 'A' {
				pts += 3
			} else if line[0] == 'C' {
				pts += 6
			}
		} else if line[2] == 'Y' {
			pts += 2
			if line[0] == 'B' {
				pts += 3
			} else if line[0] == 'A' {
				pts += 6
			}
		} else {
			pts += 3
			if line[0] == 'C' {
				pts += 3
			} else if line[0] == 'B' {
				pts += 6
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pts)
}
