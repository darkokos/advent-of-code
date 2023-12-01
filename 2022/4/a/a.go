package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	counter := 0
	for scanner.Scan() {
		line := scanner.Text()

		s_g := strings.Split(line, ",")
		g1 := s_g[0]
		g2 := s_g[1]

		g1_limit := strings.Split(g1, "-")

		g1_lower, err := strconv.Atoi(g1_limit[0])
		if err != nil {
			panic(err)
		}

		g1_upper, err := strconv.Atoi(g1_limit[1])
		if err != nil {
			panic(err)
		}

		g2_limit := strings.Split(g2, "-")

		g2_lower, err := strconv.Atoi(g2_limit[0])
		if err != nil {
			panic(err)
		}

		g2_upper, err := strconv.Atoi(g2_limit[1])
		if err != nil {
			panic(err)
		}

		if g1_lower >= g2_lower && g1_upper <= g2_upper || g2_lower >= g1_lower && g2_upper <= g1_upper {
			counter++
		}
	}

	fmt.Println(counter)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
