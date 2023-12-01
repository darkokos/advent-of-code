package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	prio_map := make(map[rune]int)
	for i := 1; i <= 26; i++ {
		prio_map[rune('a'+i-1)] = i
	}
	for i := 27; i <= 52; i++ {
		prio_map[rune('A'+i-27)] = i
	}

	f, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line)/2; i++ {
			b_c := 0
			for j := len(line) / 2; j < len(line); j++ {
				if line[i] == line[j] {
					sum += prio_map[rune(line[i])]
					b_c = 1
					break
				}
			}
			if b_c == 1 {
				break
			}
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
