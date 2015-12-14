package coal

type acolyteOfPainEffect struct {
	me Minion
}

func (this *acolyteOfPainEffect) apply(g Game, params ...interface{}) {
	if min, isMinion := params[0].(Minion); isMinion && min == this.me {
		pIdx := g.ownerOf(this.me)
		g.drawOne(pIdx)
	}
}

func newAcolyteOfPain() Minion {
	acolyteOfPain := newMinion("Acolyte of Pain", 3, 1, 3)
	acolyteOfPain.addAura(&onDamage{&acolyteOfPainEffect{acolyteOfPain}})
	return acolyteOfPain
}
