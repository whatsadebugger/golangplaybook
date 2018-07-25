package main

import (
	"bufio"
	"os"
)

// ```ruby
// letters = [
//   ["r", "h", "r", "e"],
//   ["y", "p", "c", "s"],
//   ["w", "n", "s", "n"],
//   ["t", "e", "g", "o"]
// ]
// ```

func main() {

	letters := [][]string{
		{"r", "h", "r", "e"},
		{"y", "p", "c", "s"},
		{"w", "n", "s", "n"},
		{"t", "e", "g", "o"},
	}
	f, err := os.Open("wordsEn.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	validWords := make([]string, 109583)

	i := 0

	for s.Scan() {
		validWords[i] = s.Text()
		i++
	}

	boggle(letters, validWords)
}

func boggle(letters [][]string, validWords []string) {
	type point struct {
		x int
		y int
	}

	visited := make(map[point]bool)

	_ = visited

	moves := []point{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	word := ""

	for y := range letters {
		for x := range letters[y] {
			currentSquare := point{x, y}
			visited[currentSquare] = true

			word += letters[y][x]

			for i, move := range moves {
				var possibleMove point
				possibleMove.x, possibleMove.y = currentSquare.x+move.x, currentSquare.y+move.y
				if possibleMove.x < 0 || possibleMove.x > 3 || possibleMove.y < 0 || possibleMove.y > 3 {

				} else if visited[possibleMove] == true {

				} else {
					currentSquare = possibleMove
					word += let
				}
			}

		}
	}
}
