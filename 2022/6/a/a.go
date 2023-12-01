package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func compare(s []rune) bool {
	for i, v := range s {
		for _, v1 := range s[i+1:] {
			if v == v1 {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var num int

	for scanner.Scan() {
		line := scanner.Text()

		var s []rune

		for i := 0; i < 4; i++ {
			s = append(s, rune(line[i]))
		}

		for i := 4; i < len(line); i++ {
			if !compare(s) {
				num = i
				break
			}

			s = s[1:]
			s = append(s, rune(line[i]))
		}
	}

	fmt.Println(num)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
