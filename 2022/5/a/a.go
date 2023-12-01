package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var s_l [][]string
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		if line[0][0] == '[' {

		} else if line[0] == "make" {

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
