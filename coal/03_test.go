package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test03(t *testing.T) {
	a := assert.New(t)

	g := newGame([2]*player{
		&player{
			hero:    newRexxar(),
			hand:    []card{newAbusiveSergeant()},
			minions: []minion{newChillwindYeti()},
		},
		&player{
			hero: newRexxar(),
		},
	})

	g.play(0, 0, 1)
	g.attack(1, 0)

	a.Equal(0, len(g.players[0].hand))
	a.Equal(2, len(g.players[0].minions))
	a.Equal(6, g.players[0].minions[0].attack())
	a.Equal(24, g.players[1].hero.health())

	g.nextTurn()

	a.Equal(4, g.players[1].minions[0].attack())
}
