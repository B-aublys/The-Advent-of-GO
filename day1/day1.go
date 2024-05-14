package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Executing Advent of GO day 1")

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := 0
	first_index := 0
	first_number := 0
	last_index := 0
	last_number := 0

	numbers := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		first_index = 10000
		first_number = 0
		last_index = 0
		last_number = 0

		// Get the indexes of the written number
		// Check what are the indexies of written numbers first and last
		for i, number := range numbers {
			found_index := strings.Index(scanner.Text(), number)

			if found_index == -1 {
				continue
			}

			if found_index < first_index {
				first_index = found_index
				first_number = i + 1
			}

			found_index = strings.LastIndex(scanner.Text(), number)

			if found_index == -1 {
				continue
			}

			if found_index > last_index {
				last_index = found_index
				last_number = i + 1
			}
		}

		// Iterate for the first number
		for i, character := range scanner.Text() {
			if i > first_index {
				break
			}

			if character >= 49 && character <= 58 {
				first_number = int(character - 48)
				break
			}
		}

		// Iterate backwords for the second number
		for i := len(scanner.Text()) - 1; i >= 0; i-- {
			if i < last_index {
				break
			}
			character := scanner.Text()[i]
			if character >= 49 && character <= 58 {
				last_number = int(character - 48)
				break
			}
		}

		result += first_number*10 + last_number
	}

	fmt.Printf("The resulting adjustment number: %d", result)

}
