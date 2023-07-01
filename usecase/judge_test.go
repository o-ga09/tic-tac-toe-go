package usecase_test

import (
	"go-tic-tac-toe/domain"
	mock_port "go-tic-tac-toe/mock/port"
	"go-tic-tac-toe/usecase"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestIsWin(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockGameInterface(mock)
	gameservice := usecase.ProviderGameDriver(mockResult)
	got := gameservice.IsWin(domain.Board{})
	assert.Equal(t, true, got)
}

func TestDisplay(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockResult := mock_port.NewMockGameInterface(mock)
	mockResult.EXPECT().Display(&domain.Board{}).AnyTimes().Return(nil)
	gameservice := usecase.ProviderGameDriver(mockResult)
	got := gameservice.Display(&domain.Board{})
	assert.Equal(t,nil,got)
}

func TestInput(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()
	
	mockResult := mock_port.NewMockGameInterface(mock)
	mockResult.EXPECT().Input(1).AnyTimes().Return(domain.Koma{Order: 1,X: 1,Y: 2},nil)

	gameservice := usecase.ProviderGameDriver(mockResult)
	gameservice.Init()
	got, err := gameservice.Input(1)
	assert.Equal(t,domain.Koma{Order: 1,X: 1,Y: 2},got)
	assert.Equal(t,nil,err)
}