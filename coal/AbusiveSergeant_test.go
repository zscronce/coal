package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbusiveSergeant(t *testing.T) {
	a := assert.New(t)

	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			hand:    []Card{newAbusiveSergeant()},
			minions: []Minion{newChillwindYeti()},
			mana:    10,
			maxMana: 10,
		},
		&Player{
			hero:    newRexxar(),
			mana:    10,
			maxMana: 10,
		},
	})

	g.PlayFromHand(0, 0, 1)
	g.Attack(1, 0)

	a.Equal(0, len(g.Active().hand))
	a.Equal(2, len(g.Active().minions))
	a.Equal(6, g.Active().minions[0].Attack())
	a.Equal(24, g.Inactive().hero.Health())

	g.EndTurn()

	a.Equal(4, g.Inactive().minions[0].Attack())
}
