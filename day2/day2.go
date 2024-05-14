package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// So about this solutions
// I first wanted to go the scanning character by character approach
// But decided to try the Go inbuild strings library to learn how GO
// Manages Strings and all and honestly it's very interesting

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Slice of all games
	game_maxes := make([]map[string]int, 0)

	for scanner.Scan() {
		current_game := map[string]int{"red": 0, "green": 0, "blue": 0}

		c := string(scanner.Text())
		games := strings.Split(c, ":")[1]
		split_games := strings.Split(games, ";")

		for _, game := range split_games {
			split_game := strings.Split(game, ",")
			for _, hand := range split_game {
				split_hand := strings.Split(hand, " ")
				var value int
				_, err := fmt.Sscanf(split_hand[1], "%d", &value)

				if err != nil {
					log.Fatal(err)
				}

				if current_game[split_hand[2]] < value {
					current_game[string(split_hand[2])] = value
				}
			}
		}
		game_maxes = append(game_maxes, current_game)
	}

	actual_load := map[string]int{"red": 12, "green": 13, "blue": 14}
	games_that_can_happen := 0
	for index, game := range game_maxes {
		if actual_load["red"] >= game["red"] &&
			actual_load["green"] >= game["green"] &&
			actual_load["blue"] >= game["blue"] {
			games_that_can_happen += index + 1
		}
	}

	fmt.Printf("The sum of the games that can happen are: %d", games_that_can_happen)

}
