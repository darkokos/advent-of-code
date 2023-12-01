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
	var s []string
	for scanner.Scan() {
		if len(s) != 3 {
			s = append(s, scanner.Text())
		}
		if len(s) == 3 {
			fmt.Println(s[0], s[1], s[2])
			b_c := 0
			for _, i := range s[0] {
				for _, j := range s[1] {
					if i == j {
						for _, k := range s[2] {
							if j == k {
								sum += prio_map[k]
								b_c = 1
								break
							}
						}
						if b_c == 1 {
							break
						}
					}
				}
				if b_c == 1 {
					break
				}
			}
			s = nil
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
