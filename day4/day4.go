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

	// y = C + x * x_mult
	x_mult float32
	c      float32
}

func (hs *HailStone) Calculate_formula() {
	hs.x_mult = float32(hs.y_delta) / float32(hs.x_delta)
	hs.c = -float32(hs.y) + float32(hs.x)*hs.x_mult

}

func newHailStone(x, y, x_delta, y_delta int) HailStone {
	hailStone := HailStone{x: x, y: y, x_delta: x_delta, y_delta: y_delta}
	hailStone.Calculate_formula()
	return hailStone
}

func hailstone_intersection_point(hs1, hs2 HailStone) (x, y float32) {

}

func main() {
	// [] Read all the data from the file into HailStone Structs

	file, err := os.Open("short_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	hail_storms := make([]HailStone, 0)

	for scanner.Scan() {
		text := scanner.Text()
		var x, y, trash, x_delta, y_delta int
		fmt.Println(text)
		fmt.Sscanf(text, "%d, %d, %d @ %d, %d, %d", &x, &y, &trash, &x_delta, &y_delta)
		fmt.Printf("%d %d %d %d\n", x, y, x_delta, y_delta)
		hail_storms = append(hail_storms, newHailStone(x, y, x_delta, y_delta))
	}
}
