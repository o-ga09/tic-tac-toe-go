package handler

import (
	"fmt"
	"go-tic-tac-toe/di"
	"go-tic-tac-toe/domain"
)

const (
	TURN1 = 1 //先攻
	TURN2 = 2 //後攻
)

func Run() {
	var nowTRUN int
	var in domain.Koma

	game := di.InitGame()
	board :=game.Init()
	// game.Display(board)


	nowTRUN = TURN1
	for {
		switch (nowTRUN) {
		case TURN1:
			fmt.Printf("○の番です。駒を置いてください(例:1,2と入力すると、１行２列目に駒が置かれます)")
			in, _ = game.Input(TURN1)
			fmt.Printf("now 先攻 input : \n x : %d\n y : %d\n",in.X,in.Y)
			nowTRUN = TURN2
		case TURN2:
			fmt.Printf("×の番です。駒を置いてください(例:1,2と入力すると、１行２列目に駒が置かれます)")
			in, _ = game.Input(TURN2)
			fmt.Printf("now 後攻 input : \n x : %d\n y : %d\n",in.X,in.Y)
			nowTRUN = TURN1
		}

		result := game.IsWin(*board)
		if result {
			switch (in.Order) {
			case TURN1:
				fmt.Printf("先攻勝利\n")
			case TURN2:
				fmt.Printf("後攻勝利\n")
			}
			break
		}
		game.Display(board)

	}
	game.Display(board)
}