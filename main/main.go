package main

import (
	"../tictacgo"
	"fmt"
	"strings"
)

func takeTurn (g *tictacgo.Game, isx bool) {
	tictacgo.RenderBoard(*g)
	fmt.Printf("%s, move? ", g.PlayerName(isx))

	var cell uint
	_,err := fmt.Scanln(&cell)
	cell--
	err = g.Move(isx, cell)
	if err != nil {
		fmt.Println(err.Error())
		takeTurn(g,isx)
	}
}

func playGame (g *tictacgo.Game) {
	winner := g.Winner()
	if winner == "none" {
		takeTurn(g,true)
		winner = g.Winner()
		if winner == "none" {
			takeTurn(g, false)
			winner = g.Winner()
		}
	}
	if winner != "none" {
		fmt.Println("The winner is", winner)
		tictacgo.RenderBoard(*g)
	} else {
		playGame(g)
	}
}

func getName(prompt string, askxo bool) (bool, string) {
	name := prompt
	isx := askxo
	fmt.Printf("%s name: ", prompt)
	_,err := fmt.Scanln(&name)
	if err  != nil {
		fmt.Println("Got error", err.Error())
		return false, ""
	}
	for  ; askxo;  {
		xo := "X"
		fmt.Printf("do you want X or O? [X] ")
		n,err := fmt.Scanln(&xo)
		if err != nil {
	        isx = true
	        askxo = false
			fmt.Println("Got error", err.Error())
		} else if n == 1 && strings.Contains("XxOo", xo) {
			isx = strings.Contains("xX",xo)
			askxo = false
		} else {
			fmt.Println("Enter only X or O")
		}
	}
	return isx, name
}

func main() {
	isx,p1name := getName("Player 1", true)
	_,p2name := getName("Player 2", false)
	g := new(tictacgo.Game)
	if isx {
		g.SetNames(p1name, p2name)
	} else {
		g.SetNames(p2name, p1name)
	}

    playGame (g)

}
