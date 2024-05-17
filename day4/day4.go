package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type HailStone struct {
	x       int
	y       int
	x_delta int
	y_delta int
}

func main() {
	// [] Read all the data from the file into HailStone Structs

	file, err := os.Open("short_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		var x, y, trash, x_delta, y_delta int
		fmt.Println(text)
		fmt.Sscanf(text, "%d, %d, %d @ %d, %d, %d", &x, &y, &trash, &x_delta, &y_delta)
		fmt.Printf("%d %d %d %d\n", x, y, x_delta, y_delta)
	}

}
