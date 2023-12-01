package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var s_c []int
	var s_g []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			sum := 0
			for _, v := range s_c {
				sum += v
			}

			s_g = append(s_g, sum)
			s_c = nil
			continue
		}

		c_int, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		s_c = append(s_c, c_int)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(s_g)

	sum := s_g[len(s_g)-1] + s_g[len(s_g)-2] + s_g[len(s_g)-3]

	fmt.Println(sum)
}
