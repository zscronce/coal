package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlAkirTheWindlord(t *testing.T) {
	a := assert.New(t)

	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			hand:    []Card{newArcaniteReaper()},
			minions: []Minion{newChillwindYeti()},
			mana:    10,
			maxMana: 10,
		},
		&Player{
			hero:    newRexxar(),
			hand:    []Card{newAlAkirTheWindlord()},
			mana:    10,
			maxMana: 10,
		},
	})

	g.PlayFromHand(0)
	g.EndTurn()
	g.PlayFromHand(0)

	a.Equal(true, g.Active().minions[0].DivineShield())

	g.Attack(1, 0)

	a.Equal(true, g.Active().minions[0].DivineShield())

	g.Attack(1, 1)

	a.Equal(1, len(g.Active().minions))
	a.Equal(1, len(g.Inactive().minions))
	a.Equal(true, g.Active().minions[0].Charge())
	a.Equal(false, g.Active().minions[0].DivineShield())
	a.Equal(true, g.Active().minions[0].Taunt())
	a.Equal(true, g.Active().minions[0].Windfury())
	a.Equal(5, g.Active().minions[0].Health())
	a.Equal(2, g.Inactive().minions[0].Health())
	a.Equal(27, g.Inactive().hero.Health())

	g.EndTurn()
	g.Attack(0, 1)

	a.Equal(0, len(g.Inactive().minions))
	a.Equal(24, g.Active().hero.Health())
}
