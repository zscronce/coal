package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlarmOBot(t *testing.T) {
	a := assert.New(t)

	bot1 := newAlarmOBot()
	bot2 := newAlarmOBot()
	bot3 := newAlarmOBot()

	g := NewGame([2]*Player{
		&Player{
			hero: newRexxar(),
			hand: []Card{bot1, bot2, bot3},
		},
		&Player{
			hero: newRexxar(),
		},
	})

	g.PlayFromHand(0)

	g.EndTurn()
	g.EndTurn()

	bot1InHand := false
	for _, handCard := range g.Active().hand {
		if handCard == bot1.(Card) {
			bot1InHand = true
		}
	}

	a.Equal(true, bot1InHand)
	a.Equal(2, len(g.Active().hand))
	a.Equal(1, len(g.Active().minions))

	g.PlayFromHand(0)

	a.Equal(1, len(g.Active().hand))
	a.Equal(2, len(g.Active().minions))

	g.PlayFromHand(0)

	a.Equal(0, len(g.Active().hand))
	a.Equal(3, len(g.Active().minions))

	g.EndTurn()
	g.EndTurn()

	a.Equal(0, len(g.Active().hand))
	a.Equal(3, len(g.Active().minions))
}
