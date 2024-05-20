package main

import (
	"bufio"
	"errors"
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
	hs.c = float32(hs.y) - float32(hs.x)*hs.x_mult
}

func newHailStone(x, y, x_delta, y_delta int) *HailStone {
	hailStone := &HailStone{x: x, y: y, x_delta: x_delta, y_delta: y_delta}
	hailStone.Calculate_formula()
	return hailStone
}

// TODO: add a test if they have touched in the past :D
func hailstone_intersection_point(hs1, hs2 *HailStone) (x, y float32, err error) {

	if hs1.x_mult-hs2.x_mult == 0 {
		return 0, 0, errors.New("lines are parallel")
	}

	intersection_x := -(hs1.c - hs2.c) / (hs1.x_mult - hs2.x_mult)

	// Checks top check of the rocks have intersected in the past
	if (intersection_x < float32(hs1.x) && hs1.x_delta > 0) ||
		(intersection_x > float32(hs1.x) && hs1.x_delta < 0) {
		return 0, 0, errors.New("the stones hve touched it the past")
	}

	if (intersection_x < float32(hs2.x) && hs2.x_delta > 0) ||
		(intersection_x > float32(hs2.x) && hs2.x_delta < 0) {
		return 0, 0, errors.New("the stones hve touched it the past")
	}

	intersection_y := float32(hs1.c) + (intersection_x * hs1.x_mult)

	return intersection_x, intersection_y, nil

}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	hail_storms := make([]*HailStone, 0)

	for scanner.Scan() {
		text := scanner.Text()
		var x, y, trash, x_delta, y_delta int
		fmt.Sscanf(text, "%d, %d, %d @ %d, %d, %d", &x, &y, &trash, &x_delta, &y_delta)
		hail_storms = append(hail_storms, newHailStone(x, y, x_delta, y_delta))
	}

	sum_in_area := 0

	for i := 0; i < len(hail_storms)-1; i++ {
		for j := i + 1; j < len(hail_storms); j++ {
			x, y, err := hailstone_intersection_point(hail_storms[i], hail_storms[j])
			if x >= 200000000000000 && x <= 400000000000000 && y >= 200000000000000 && y <= 400000000000000 && err == nil {
				fmt.Println("made it: ")
				fmt.Printf("%d:%d, [ %.3f | %.3f ]\n", i, j, x, y)
				sum_in_area += 1
			}
		}
	}

	fmt.Printf("Amount of hailstorms in the colltion area: %d", sum_in_area)
}
