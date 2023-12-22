package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	row int
	col int
}

func main() {
	clearScreen()
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("------------------------------")

Loop:
	for {
		clearScreen()
		menuChoice := displayMenu()

		switch menuChoice {
		case "s":
			runGame()
		case "q":
			break Loop
		default:
			clearScreen()
			fmt.Println("\nInvalid choice. Please enter a valid option.\n")
		}

	}
	clearScreen()
	fmt.Println("Thanks for Playing!")

}

func displayMenu() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Menu Options: ")
	fmt.Println("- start game: s --------------")
	fmt.Println("- quit game:  q --------------")
	fmt.Println("------------------------------")
	fmt.Println("enter: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func clearScreen() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("------------------------------")
}

func runGame() {
	reader := bufio.NewReader(os.Stdin)

	board := [3][3]interface{}{
		{1, 2, 3}, /*  initializers for row indexed by 0 */
		{4, 5, 6}, /*  initializers for row indexed by 1 */
		{7, 8, 9}, /*  initializers for row indexed by 2 */
	}

	var moves []int
	gameIsWon := false
	playerIsX := true

Game:
	for {
		displayBoard(board)
		fmt.Println("Choose a spot on the board: ")
		playerInput, _ := reader.ReadString('\n')
		playerInput = strings.Replace(playerInput, "\n", "", -1)
		validMove := isValidMove(playerInput, moves)
		switch validMove {
		case true:
			move, _ := strconv.Atoi(playerInput)
			board = insertMove(move, playerIsX, board)
			moves = append(moves, move)

			gameIsWon = validateWin(board, move)

			if gameIsWon || len(moves) == 9 {
				break Game
			} else {
				playerIsX = !playerIsX
			}
			break
		default:
			clearScreen()
			fmt.Println("\nInvalid choice. Please enter a valid option.\n")
		}
	}

	displayBoard(board)
	if gameIsWon {
		var player string
		if playerIsX {
			player = "1"
		} else {
			player = "2"
		}
		fmt.Printf("\nPlayer %v is the winner!\n", player)
	} else {
		fmt.Println("\nGame is a tie!")
	}

}

func displayBoard(board [3][3]interface{}) {
	clearScreen()
	fmt.Printf(" %v | %v | %v \n", board[0][0], board[0][1], board[0][2])
	fmt.Println("--- --- ---")
	fmt.Printf(" %v | %v | %v \n", board[1][0], board[1][1], board[1][2])
	fmt.Println("--- --- ---")
	fmt.Printf(" %v | %v | %v \n", board[2][0], board[2][1], board[2][2])
}

func isValidMove(newMove string, moves []int) bool {
	move, err := strconv.Atoi(newMove)

	if err != nil {
		return false
	}

	for _, played := range moves {
		if played == move {
			return false
		}
	}

	return true
}

func insertMove(move int, playerIsX bool, board [3][3]interface{}) [3][3]interface{} {
	coord := moveToRowCol(move)
	if playerIsX {
		board[coord.row][coord.col] = "x"

	} else {
		board[coord.row][coord.col] = "o"

	}
	return board
}

func moveToRowCol(move int) Coord {
	rows := [3]int{1, 4, 7}
	var row int
	if 1 <= move && move <= 3 {
		row = 0
	} else if 4 <= move && move <= 6 {
		row = 1
	} else {
		row = 2
	}
	col := move - rows[row]
	return Coord{row, col}
}

func validateWin(board [3][3]interface{}, lastMove int) bool {
	coord := moveToRowCol(lastMove)
	row, col := coord.row, coord.col

	if (board[row][col] == board[row][(col+1)%3]) && (board[row][col] == board[row][(col+2)%3]) {
		return true
	}

	if (board[row][col] == board[(row+1)%3][col]) && (board[row][col] == board[(row+2)%3][col]) {
		return true
	}

	switch lastMove {
	case 1, 9:
		if (board[0][0] == board[1][1]) && (board[1][1] == board[2][2]) {
			return true
		}
	case 3, 7:
		if (board[0][2] == board[1][1]) && (board[1][1] == board[2][0]) {
			return true
		}
	case 5:
		if (board[0][0] == board[1][1]) && (board[1][1] == board[2][2]) || (board[0][2] == board[1][1]) && (board[1][1] == board[2][0]) {
			return true
		}
	default:
		break
	}
	return false
}
