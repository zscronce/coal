package coal

type destroyOtherMinion struct {
	me Minion
}

func (this *destroyOtherMinion) apply(g Game, params ...interface{}) {
	if min, isMinion := params[0].(Minion); isMinion && min != this.me {
		g.destroy(min)
	}
}

func newAcidmaw() Minion {
	acidmaw := newMinion("Acidmaw", 7, 4, 2)
	acidmaw.addAura(&onDamage{&destroyOtherMinion{acidmaw}})
	return acidmaw
}
