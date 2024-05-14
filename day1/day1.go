package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Executing Advent of GO day 1")

	file, err := os.Open("short_input.txt")

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
		fmt.Println(scanner.Text())

		for _, character := range scanner.Text() {
			if character >= 49 && character <= 58 {
				adjusted_number := character - 48
				if first_number == 0 {
					first_number = int(adjusted_number)
					last_number = int(adjusted_number)
				} else {
					last_number = int(adjusted_number)
				}
			}
		}
		fmt.Println(first_number*10 + last_number)
		result += first_number*10 + last_number
	}

	fmt.Printf("The resulting adjustment number: %d", result)

}
