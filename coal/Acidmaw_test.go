package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcidmaw(t *testing.T) {
	a := assert.New(t)

	g := NewGame([2]*Player{
		&Player{
			hero:    newRexxar(),
			minions: []Minion{newChillwindYeti(), newAcidmaw()},
		},
		&Player{
			hero:    newRexxar(),
			minions: []Minion{newAcidmaw(), newAcidmaw()},
		},
	})

	g.damageAllCharacters(1)

	a.Equal(0, len(g.Players()[0].minions))
	a.Equal(0, len(g.Players()[1].minions))
}
