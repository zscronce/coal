package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArcaniteReaper(t *testing.T) {
	a := assert.New(t)

	g := newGame([2]*player{
		&player{
			hero: newRexxar(),
			hand: []card{newArcaniteReaper()},
		},
		&player{
			hero: newRexxar(),
			hand: []card{newArcaniteReaper()},
		},
	})

	g.play(0)
	g.nextTurn()
	g.play(0)
	g.attack(0, 0)

	a.Equal(5, g.players[0].hero.weapon().attack())
	a.Equal(5, g.players[1].hero.weapon().attack())
	a.Equal(1, g.players[0].hero.weapon().durability())
	a.Equal(2, g.players[1].hero.weapon().durability())
	a.Equal(30, g.players[0].hero.health())
	a.Equal(25, g.players[1].hero.health())

	g.attack(0, 0)

	a.Equal(20, g.players[1].hero.health())
	a.Equal(30, g.players[0].hero.health())
	a.Equal(nil, g.players[0].hero.weapon())

	g.attack(0, 0)

	a.Equal(20, g.players[1].hero.health())
	a.Equal(30, g.players[0].hero.health())
	a.Equal(nil, g.players[0].hero.weapon())
}
