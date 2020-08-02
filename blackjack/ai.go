package blackjack

import (
	"fmt"

	"github.com/yuanyu90221/deck"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

func HumanAI() AI {
	return humanAI{}
}

type humanAI struct{}

type dealerAI struct{}

func (ai dealerAI) Bet() int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || (dScore == 17 && Soft(hand...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	//
}

func (ai humanAI) Bet() int {
	return 1
}
func (ai humanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hand)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		default:
			fmt.Println("Invalid Option:", input)
		}
	}
}

type GameState struct {
}

// type Move func(GameState) GameState

// func Hit(gs GameState) GameState {
// 	return gs
// }

// func Stand(gs GameState) GameState {
// 	return gs
// }

func (ai humanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS===")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)
}
