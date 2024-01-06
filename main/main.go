package main

import (
	"../tictacgo"
	"fmt"
)

func takeTurn (p1, p2 *tictacgo.Player) {
	tictacgo.RenderBoard(*p1, *p2)
	fmt.Printf("%s, move? ", p1.Name)
	var cell uint
	_,err := fmt.Scanln(&cell)
	cell--
	err = tictacgo.Move(p1, p2, cell)
	if err != nil {
		fmt.Println(err.Error())
		takeTurn(p1,p2)
	}
}

func playGame (p1, p2 *tictacgo.Player) {
	winner := tictacgo.Winner(p1, p2)
	if winner == "none" {
		takeTurn(p1,p2)
		winner = tictacgo.Winner(p1, p2)
		if winner == "none" {
			takeTurn(p2, p1)
			winner = tictacgo.Winner(p2, p1)
		}
	}
	if winner != "none" {
		fmt.Println("The winner is", winner)
		tictacgo.RenderBoard(*p1,*p2)
	} else {
		playGame(p1,p2)
	}
}

func main() {
	tictacgo.InitWins()
	var name string
	var p1, p2 tictacgo.Player
	fmt.Printf("Player 1 name: ")
	_,err := fmt.Scanln(&name)
	if err  != nil {
		fmt.Println("Got error", err.Error())
		return
	}
	tictacgo.InitPlayer(&p1, name, true)
	fmt.Printf("Player 2 name: ")
	_,err = fmt.Scanln(&name)
	if err  != nil {
		fmt.Println("Got error", err.Error())
		return
	}
	tictacgo.InitPlayer(&p2, name, false)
    playGame (&p1, &p2)

}
