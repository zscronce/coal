package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcolyteOfPain(t *testing.T) {
	a := assert.New(t)

	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			hand:    []Card{newAcolyteOfPain()},
			deck:    []Card{newAcolyteOfPain(), newAcolyteOfPain()},
			mana:    10,
			maxMana: 10,
		},
		&Player{
			hero: newRexxar(),
			deck: []Card{newAcolyteOfPain(), newAcolyteOfPain()},
		},
	})

	g.PlayFromHand(0)
	g.damageAllCharacters(1)

	a.Equal(2, g.Players()[0].minions[0].Health())
	a.Equal(1, len(g.Players()[0].hand))
	a.Equal(1, len(g.Players()[0].deck))
	a.Equal(0, len(g.Players()[1].hand))
	a.Equal(2, len(g.Players()[1].deck))

	g.EndTurn()
	g.damageAllCharacters(1)

	a.Equal(1, g.Players()[1].minions[0].Health())
	a.Equal(2, len(g.Players()[1].hand))
	a.Equal(0, len(g.Players()[1].deck))
	a.Equal(0, len(g.Players()[0].hand))
	a.Equal(2, len(g.Players()[0].deck))

	g.damageAllCharacters(1)

	a.Equal(0, len(g.Players()[1].minions))
	a.Equal(2, len(g.Players()[1].hand))
	a.Equal(0, len(g.Players()[1].deck))
	a.Equal(26, g.Players()[1].hero.Health())
	a.Equal(0, len(g.Players()[0].hand))
	a.Equal(2, len(g.Players()[0].deck))
}
