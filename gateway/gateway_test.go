package gateway_test

import (
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/domain"
	"go-tic-tac-toe/gateway"
	mock_ui "go-tic-tac-toe/mock/ui"
	"io"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestInput(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	var in io.Reader
	mockResult := mock_ui.NewMockIODriver(mock)
	mockResult.EXPECT().Input(in).AnyTimes().Return(ui.DriverKoma{X: 0,Y: 0},nil)
	gamegateway := gateway.ProviderGameDriver(mockResult)
	got, err := gamegateway.Input(0)
	assert.Equal(t, domain.Koma{Order: 0,X: 0,Y: 0}, got)
	assert.Equal(t,nil,err)
}

func TestDisplay(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_ui.NewMockIODriver(mock)
	mockResult.EXPECT().Display(ui.DriverBoard{
		Board: [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}},
	}).AnyTimes().Return(nil)
	gamegateway := gateway.ProviderGameDriver(mockResult)
	err := gamegateway.Display(&domain.Board{})
	assert.Equal(t,nil,err)
}