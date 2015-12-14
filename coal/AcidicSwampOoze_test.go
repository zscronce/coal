package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcidicSwampOoze(t *testing.T) {
	a := assert.New(t)

	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			hand:    []Card{newArcaniteReaper()},
			mana:    10,
			maxMana: 10,
		},
		&Player{
			hero:    newRexxar(),
			hand:    []Card{newAcidicSwampOoze()},
			mana:    10,
			maxMana: 10,
		},
	})

	g.PlayFromHand(0)
	g.Attack(0, 0)

	a.Equal(newArcaniteReaper().Name(), g.Players()[0].hero.Weapon().Name())
	a.Equal(25, g.Players()[1].hero.Health())

	g.EndTurn()
	g.PlayFromHand(0)

	a.Equal(1, len(g.Players()[0].minions))
	a.Equal(nil, g.Players()[1].hero.Weapon())
}
