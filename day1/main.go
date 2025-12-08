package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func parseInput(inputs []string) []int {
	var actions []int

	for i := range inputs {
		input := inputs[i]

		distance, err := strconv.Atoi(input[1:])
		if err != nil {
			log.Fatalf("couldn't parse input: %s", err)
		}
		if input[0:1] == "L" {
			distance *= -1
		}
		actions = append(actions, distance)
	}
	return actions
}

func resolveActions(actions []int, position int) int {
	password := 0
	position += actions[0] % 100
	if position > 99 {
		position -= 100
	} else if position < 0 {
		position += 100
	}

	if position == 0 {
		password++
	}
	if len(actions[1:]) > 0 {
		password += resolveActions(actions[1:], position)
	}
	return password
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()
	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	password := resolveActions(parseInput(inputs), 50)
	log.Printf("Password: %o", password)
}
