package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArcaniteReaper(t *testing.T) {
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
			hand:    []Card{newArcaniteReaper()},
			mana:    10,
			maxMana: 10,
		},
	})

	g.PlayFromHand(0)
	g.EndTurn()
	g.PlayFromHand(0)
	g.Attack(0, 0)

	a.Equal(5, g.Active().hero.Weapon().Attack())
	a.Equal(5, g.Inactive().hero.Weapon().Attack())
	a.Equal(1, g.Active().hero.Weapon().Durability())
	a.Equal(2, g.Inactive().hero.Weapon().Durability())
	a.Equal(30, g.Active().hero.Health())
	a.Equal(25, g.Inactive().hero.Health())

	g.Attack(0, 0)

	a.Equal(20, g.Inactive().hero.Health())
	a.Equal(30, g.Active().hero.Health())
	a.Equal(nil, g.Active().hero.Weapon())

	g.Attack(0, 0)

	a.Equal(20, g.Inactive().hero.Health())
	a.Equal(30, g.Active().hero.Health())
	a.Equal(nil, g.Active().hero.Weapon())
}
