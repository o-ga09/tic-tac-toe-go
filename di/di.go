package di

import (
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/gateway"
	"go-tic-tac-toe/usecase"
)

func InitGame() *usecase.GameService {
	ui := ui.PrioviderUIDriver()
	gateway := gateway.ProviderGameDriver(ui)
	usecase := usecase.ProviderGameDriver(gateway)

	return &usecase
}