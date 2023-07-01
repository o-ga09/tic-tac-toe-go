package gateway

import (
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/domain"
	"io"
)

type GameGateway struct {
	gameDriver ui.IODriver
}

func ProviderGameDriver(gameDriver ui.IODriver) GameGateway {
	return GameGateway{gameDriver: gameDriver}
}

func (gate *GameGateway) Input() (domain.Koma, error) {
	var in io.Reader
	res, err := gate.gameDriver.Input(in)
	return domain.Koma{Order: res.Order,X: res.X,Y: res.Y}, err
}

func (gate *GameGateway) Display(board *domain.Board) error {
	err := gate.gameDriver.Display(ui.DriverBoard{})
	return err
}

func (gate *GameGateway) IsWin(board domain.Board) bool {
	return true
}