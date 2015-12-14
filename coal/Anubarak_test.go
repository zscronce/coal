package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnubarak(t *testing.T) {
	a := assert.New(t)

	anubarak := newAnubarak()
	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			hand:    []Card{anubarak},
			mana:    10,
			maxMana: 10,
		},
		&Player{
			hero:    newRexxar(),
			minions: []Minion{newChillwindYeti()},
			mana:    10,
			maxMana: 10,
		},
	})

	g.PlayFromHand(0)

	a.Equal(1, g.Active().mana)
	a.Equal(0, len(g.Active().hand))
	a.Equal(1, len(g.Active().minions))
	a.Equal(anubarak, g.Active().minions[0])

	g.EndTurn()
	g.Attack(1, 1)

	a.Equal(0, len(g.Active().minions))
	a.Equal(1, len(g.Inactive().minions))
	a.Equal("Nerubian", g.Inactive().minions[0].Name())
	a.Equal(4, g.Inactive().minions[0].Attack())
	a.Equal(4, g.Inactive().minions[0].Health())
	a.Equal(4, g.Inactive().minions[0].MaxHealth())
}
