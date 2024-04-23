package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/snake-ladder/models"
)

func askNumber(message string) int {
	var num int
	fmt.Print(message)
	fmt.Scan(&num)
	bufio.NewReader(os.Stdin).ReadString('\n')
	return num
}

// TODO: change function name to askSnakesNumber
func askNumSnake() int {
	return askNumber("Type snake number: ")
}

// TODO: change function name to askLadderNumber
func askNumLadder() int {
	return askNumber("Type ladder number: ")
}

func main() {
	// TODO: delete this line and change numSnakes => snakeNumber, numLadders => ladderNumber
	numSnakes := askNumSnake()
	numLadders := askNumLadder()
	// TODO: default value for size of board
	game := models.NewGame(numSnakes, numLadders, 10)
	game.AddPlayer("red")
	game.AddPlayer("green")
	game.AddPlayer("blue")
	game.Play()
}
