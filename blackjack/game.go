package blackjack

import (
	"fmt"

	"github.com/yuanyu90221/deck"
)

const (
	statePlayerTurn state = iota
	stateDealerTurn
	stateHandOver
)

type state int8

func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAI: dealerAI{},
		balance:  0,
	}
}

type Game struct {
	// upexpeorted fields
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI AI
	balance  int
}

func deal(g *Game) {
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.deck = draw(g.deck)
		g.dealer = append(g.dealer, card)
	}
	g.state = statePlayerTurn
}

func minScore(hand ...deck.Card) int {
	score := 0
	for _, c := range hand {
		score += min(int(c.Rank), 10)
	}
	return score
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}
	for _, c := range hand {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, ane we are changing it to be worth 11
			return minScore + 10
		}
	}
	return minScore
}

func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand...)
	return minScore != score
}

func (g *Game) Play(ai AI) int {
	g.deck = deck.New(deck.Deck(3), deck.Shutfle)

	for i := 0; i < 2; i++ {
		deal(g)
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(g.player, g.dealer[0])
			move(g)
		}

		for g.state == stateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAI.Play(hand, g.dealer[0])
			move(g)
		}
		endHand(g, ai)
	}
	return g.balance
}

func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	//TODO
	switch {
	case pScore > 21:
		fmt.Println("You busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You win!")
		g.balance++
	case dScore > pScore:
		fmt.Println("You lose!")
		g.balance--
	case dScore == pScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}

type Move func(*Game)

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't currently any player's turn")
	}
}
func MoveHit(g *Game) {
	hand := g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if Score(*hand...) > 21 {
		MoveStand(g)
	}
}

func MoveStand(g *Game) {
	g.state++
}

func clone(gs Game) Game {
	ret := Game{
		deck:   make([]deck.Card, len(gs.deck)),
		state:  gs.state,
		player: make([]deck.Card, len(gs.player)),
		dealer: make([]deck.Card, len(gs.dealer)),
	}
	copy(ret.deck, gs.deck)
	copy(ret.player, gs.player)
	copy(ret.dealer, gs.dealer)
	return ret
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
