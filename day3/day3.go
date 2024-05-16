package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

// The main idea behind this solution is as follows:
// we need to think about this problem from the perspective of numbers.
// Because each number can added only once, meaning that if there are multiple
// Symbols around the number we might end including a few numebrs twice or 3 times
// Meaning that we would have to track what numebers are included and what aren't
// Which I think would be quite complicated

type number_struct struct {
	value       int
	start_index int
	end_index   int
}

func (nb number_struct) String() string {
	return fmt.Sprintf("[%d:%d] -- %d", nb.start_index, nb.end_index, nb.value)
}

func (nb number_struct) to_keep(first_line, second_line, third_line string) bool {

	var start_index int
	var end_index int

	if nb.start_index != 0 {
		start_index = nb.start_index - 1
	} else {
		start_index = nb.start_index
	}

	if nb.end_index != len(first_line)-1 {
		end_index = nb.end_index + 1
	} else {
		end_index = nb.end_index
	}

	fmt.Printf("Start: %d, End: %d", start_index, end_index)

	for i := start_index; i <= end_index; i++ {
		if is_special_character(first_line[i]) || is_special_character(third_line[i]) {
			return true
		}

		if is_special_character(second_line[start_index]) || is_special_character(second_line[end_index]) {
			return true
		}

	}
	return false
}

func is_special_character(a byte) bool {
	return !unicode.IsDigit(rune(a)) && string(a) != "."
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var last_line string
	var first_line string
	var middle_line string
	var third_line string

	// Init the first line
	scanner.Scan()
	middle_line = scanner.Text()

	for range middle_line {
		last_line += "."
	}
	first_line = last_line

	result := 0

	last_line_correction := true
	for last_line_correction {

		if !scanner.Scan() {
			third_line = last_line
			last_line_correction = false
		} else {
			third_line = scanner.Text()
		}

		numbers := make([]*number_struct, 0)
		var current_number *number_struct

		for index, character := range middle_line {
			if character > 47 && character < 58 {
				fmt.Println(string(character))
				if current_number == nil {
					current_number = &number_struct{value: 0, start_index: index, end_index: index}
				}
				current_number.value = current_number.value*10 + int(character) - 48
				current_number.end_index = index

			} else {
				if current_number != nil {
					numbers = append(numbers, current_number)
					current_number = nil
				}
			}
		}
		// Some numbers are on the edge and the loop terminates before it can check
		// if the number has ended, but we know it has, because the loop ended
		if current_number != nil {
			numbers = append(numbers, current_number)
		}

		for _, number := range numbers {
			if number.to_keep(first_line, middle_line, third_line) {
				result += number.value
				fmt.Printf("Number selected: %d\n", number.value)
			}
		}

		first_line, middle_line = middle_line, third_line
	}

	fmt.Printf("the result of day 3 is: %d\n", result)
}
