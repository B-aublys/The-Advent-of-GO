package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	first_number := 0
	last_number := 0

	for scanner.Scan() {
		first_number = 0
		last_number = 0

		// This can be done with one for loop
		// But is a bit more efficient if we loop backwards for the
		// Second number

		// Iterate for the first number
		for _, character := range scanner.Text() {
			if character >= 49 && character <= 58 {
				first_number = int(character - 48)
				break
			}
		}

		// Iterate backwords for the second number
		for i := len(scanner.Text()) - 1; i >= 0; i-- {
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
