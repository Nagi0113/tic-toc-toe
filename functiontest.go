package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	board     [3][3]string
	player    string
	gameOver  bool
	winner    string
	moveCount int
)

func main() {
	initBoard()
	printBoard()

	reader := bufio.NewReader(os.Stdin)

	for !gameOver {
		fmt.Printf("TURN OF PLAYER %s ，PLEASE INPUT（LIKE：1,1）：", player)

		input, _ := reader.ReadString('\n')
		input = strings.Trim(input, "\n")

		if input == "q" {
			fmt.Println("GAME OVER.")
			return
		}

		if !makeMove(input) {
			fmt.Println("INVALID MOVE，PLEASE RETRY")
		}

		printBoard()
		checkGameOver()
		switchPlayer()
	}

	if winner != "" {
		fmt.Printf("PLAYER %s WON！\n", winner)
	} else {
		fmt.Println("DRAW")
	}
}

func initBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "."
		}
	}
	player = "X"
	gameOver = false
	winner = ""
	moveCount = 0
}

func printBoard() {
	for i := 0; i < 3; i++ {
		fmt.Printf(" %s  %s  %s \n", board[i][0], board[i][1], board[i][2])
	}
}

func makeMove(input string) bool {
	coords := strings.Split(input, ",")
	if len(coords) != 2 {
		return false
	}

	row, col := parseMove(coords[0], coords[1])
	if row == -1 || col == -1 {
		return false
	}

	if board[row][col] != "." {
		return false
	}

	board[row][col] = player
	moveCount++
	return true
}

func parseMove(rowInput, colInput string) (int, int) {
	row, col := -1, -1

	switch rowInput {
	case "0":
		row = 0
	case "1":
		row = 1
	case "2":
		row = 2
	}

	switch colInput {
	case "0":
		col = 0
	case "1":
		col = 1
	case "2":
		col = 2
	}

	return row, col
}

func checkGameOver() {
	// CHECK ROW
	for i := 0; i < 3; i++ {
		if board[i][0] != "." && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			gameOver = true
			winner = board[i][0]
			return
		}
	}

	// CHECK COLUMN
	for i := 0; i < 3; i++ {
		if board[0][i] != "." && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			gameOver = true
			winner = board[0][i]
			return
		}
	}

	// CHECK DIAGONAL
	if board[0][0] != "." && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		gameOver = true
		winner = board[0][0]
		return
	}

	if board[0][2] != "." && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		gameOver = true
		winner = board[0][2]
		return
	}

	// DRAW CHECK
	if moveCount == 9 {
		gameOver = true
	}
}

func switchPlayer() {
	if player == "X" {
		player = "O"
	} else {
		player = "X"
	}
}
