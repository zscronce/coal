package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test02(t *testing.T) {
	a := assert.New(t)

	anubarak := newAnubarak()
	g := newGame([2]*player{
		&player{
			hero: newRexxar(),
			hand: []card{anubarak},
			mana: 10,
		},
		&player{
			hero:    newRexxar(),
			minions: []minion{newChillwindYeti()},
		},
	})

	g.play(0)

	a.Equal(1, g.players[0].mana)
	a.Equal(0, len(g.players[0].hand))
	a.Equal(1, len(g.players[0].minions))
	a.Equal(anubarak, g.players[0].minions[0])

	g.nextTurn()
	g.attack(1, 1)

	a.Equal(0, len(g.players[0].minions))
	a.Equal(1, len(g.players[1].minions))
	a.Equal("Nerubian", g.players[1].minions[0].name())
	a.Equal(4, g.players[1].minions[0].attack())
	a.Equal(4, g.players[1].minions[0].health())
	a.Equal(4, g.players[1].minions[0].maxHealth())
}
