package main

import "fmt"

type CellValue int

const (
	Empty CellValue = 0
	X     CellValue = 1
	O     CellValue = 2
)

type Cell struct {
	Value CellValue
}

type Board struct {
	cells [3][3]Cell
}

type Player struct {
	Symbol CellValue
}

type GameState int

const (
	Playing GameState = 0
	XWins   GameState = 1
	OWins   GameState = 2
	Draw    GameState = 3
)

type Game struct {
	board         Board
	currentPlayer Player
	gameState     GameState
}

func (g *Game) checkWin() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.board.cells[i][0].Value != Empty && g.board.cells[i][0].Value == g.board.cells[i][1].Value && g.board.cells[i][1].Value == g.board.cells[i][2].Value {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if g.board.cells[0][i].Value != Empty && g.board.cells[0][i].Value == g.board.cells[1][i].Value && g.board.cells[1][i].Value == g.board.cells[2][i].Value {
			return true
		}
	}

	// Check diagonals
	if g.board.cells[0][0].Value != Empty && g.board.cells[0][0].Value == g.board.cells[1][1].Value && g.board.cells[1][1].Value == g.board.cells[2][2].Value {
		return true
	}
	if g.board.cells[0][2].Value != Empty && g.board.cells[0][2].Value == g.board.cells[1][1].Value && g.board.cells[1][1].Value == g.board.cells[2][0].Value {
		return true
	}

	return false
}

func (g *Game) Play(row, col int) error {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return fmt.Errorf("invalid move")
	}
	if g.board.cells[row][col].Value != Empty {
		return fmt.Errorf("cell already occupied")
	}

	g.board.cells[row][col].Value = g.currentPlayer.Symbol

	if g.checkWin() {
		g.gameState = XWins
		if g.currentPlayer.Symbol == O {
			g.gameState = OWins
		}
		return nil
	}

	if g.checkDraw() {
		g.gameState = Draw
		return nil
	}

	if g.currentPlayer.Symbol == X {
		g.currentPlayer.Symbol = O
	} else {
		g.currentPlayer.Symbol = X
	}

	return nil
}

func (g *Game) checkDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.board.cells[i][j].Value == Empty {
				return false
			}
		}
	}
	return true
}

func (g *Game) PrintBoard() {
	for _, row := range g.board.cells {
		for _, cell := range row {
			switch cell.Value {
			case Empty:
				fmt.Print("|_._|")
			case X:
				fmt.Print("|_X_|")
			case O:
				fmt.Print("|_O_|")
			}
		}
		fmt.Println()
	}
}

func main() {
	game := Game{
		currentPlayer: Player{Symbol: X},
	}

	for game.gameState == Playing {
		game.PrintBoard()
		fmt.Println("Player", game.currentPlayer.Symbol, "turn")

		var row, col int
		fmt.Print("Enter row (0-2): ")
		fmt.Scan(&row)
		fmt.Print("Enter column (0-2): ")
		fmt.Scan(&col)

		err := game.Play(row, col)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	game.PrintBoard()
	switch game.gameState {
	case XWins:
		fmt.Println("X wins!")
	case OWins:
		fmt.Println("O wins!")
	case Draw:
		fmt.Println("Draw!")
	}
}
