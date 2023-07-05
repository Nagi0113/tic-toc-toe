package main
import(
    "fmt"
    "strings"
    "os"
    )
    var(
        board   [3][3]string
        )
    
    func main(){
        Printboard()
        Checkgameover()
    }
    
    func Printboard(){
        for i := 0; i < 3; i++{
            fmt.Printf(" %s %s %s \n",board[i][0], board[i][1], board[i][2])
        }
    }
func checkGameOver() {
	// CHECK ROW
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			gameOver = true
			winner = board[i][0]
			return
		}
	}

	// CHECK COLUMN
	for i := 0; i < 3; i++ {
		if board[0][i] != " " && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			gameOver = true
			winner = board[0][i]
			return
		}
	}

	// CHECK DIAGONAL
	if board[0][0] != " " && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		gameOver = true
		winner = board[0][0]
		return
	    }

	if board[0][2] != " " && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		gameOver = true
		winner = board[0][2]
		return
	    }

	// DRAW CHECK
	if moveCount == 9 {
		gameOver = true
	    }
    }
	
    