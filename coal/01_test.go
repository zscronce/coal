package coal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test01(t *testing.T) {
	g := &game{
		players: [2]*player{
			&player{
				hero:    newRexxar(),
				minions: []minion{newChillwindYeti()},
			},
			&player{
				hero:    newRexxar(),
				minions: []minion{newAbomination(), newChillwindYeti()},
			},
		},
	}

	g.attack(1, 1)

	characters := g.characters()
	a := assert.New(t)
	a.Equal(1, len(characters[0]))
	a.Equal(28, characters[0][0].health())
	a.Equal(2, len(characters[1]))
	a.Equal(28, characters[1][0].health())
	a.Equal(3, characters[1][1].health())
}
