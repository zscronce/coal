package coal

type effect interface {
	apply(*game, ...int)
}

type damageAllCharacters struct {
	damage int
}

type attackMod struct {
	delta *delta
}

type tempAttackMod struct {
	attackMod
}

func (this *damageAllCharacters) apply(g *game, params ...int) {
	damages := []damage{}

	for p := range g.players {
		for c := range g.players[p].characters() {
			damages = append(damages, damage{p, c, this.damage})
		}
	}

	g.damage(damages)
}

func newTempAttackDelta(a int) effect {
	return &tempAttackMod{
		attackMod: attackMod{
			delta: &delta{
				func(attack int) int {
					return attack + a
				},
			},
		},
	}
}

func (this *attackMod) getCharacter(g *game, params ...int) character {
	return g.characters()[params[0]][params[1]]
}

func (this *attackMod) applyTo(char character) {
	char.changeAttack(this.delta)
}

func (this *attackMod) apply(g *game, params ...int) {
	this.applyTo(this.getCharacter(g, params...))
}

func (this *tempAttackMod) apply(g *game, params ...int) {
	char := this.getCharacter(g, params...)
	this.applyTo(char)
	g.onceNextTurn_ = append(g.onceNextTurn_, func() {
		char.unChangeAttack(this.delta)
	})
}
