package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbomination(t *testing.T) {
	a := assert.New(t)

	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			minions: []Minion{newChillwindYeti()},
			mana:    10,
			maxMana: 10,
		},
		&Player{
			hero:    newRexxar(),
			minions: []Minion{newAbomination(), newChillwindYeti()},
			mana:    10,
			maxMana: 10,
		},
	})

	g.Attack(1, 1)

	a.Equal(0, len(g.Active().minions))
	a.Equal(28, g.Active().hero.Health())
	a.Equal(1, len(g.Inactive().minions))
	a.Equal(28, g.Inactive().hero.Health())
	a.Equal(3, g.Inactive().minions[0].Health())
}
