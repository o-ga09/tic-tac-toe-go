package gateway

import (
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/domain"
	"io"
)

type GameGateway struct {
	gameDriver ui.IODriver
}

func ProviderGameDriver(gameDriver ui.IODriver) *GameGateway {
	return &GameGateway{gameDriver: gameDriver}
}

func (gate *GameGateway) Input(player int) (domain.Koma, error) {
	var in io.Reader
	res, err := gate.gameDriver.Input(in)
	return domain.Koma{Order: player,X: res.X,Y: res.Y}, err
}

func (gate *GameGateway) Display(board *domain.Board) error {
	// 2次元スライスへの変換
	slice := make([][]string, len(board.Board))
	for i := range board.Board {
		slice[i] = make([]string, len(board.Board[i]))
		copy(slice[i], board.Board[i][:])
	}
	driverBoard := ui.DriverBoard{
		Board: slice,
	}

	err := gate.gameDriver.Display(driverBoard)
	return err
}

func (gate *GameGateway) IsWin(board domain.Board) bool {
	return true
}