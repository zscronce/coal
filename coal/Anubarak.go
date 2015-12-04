package coal

type anubarakDeathrattle struct {
	me minion
}

func (this *anubarakDeathrattle) apply(g *game, params ...int) {
	owner := g.owner[this.me]
	g.addToHand(owner, this.me)
	g.summon(owner, newMinion("Nerubian", 2, 4, 4))
}

func newAnubarak() minion {
	anubarak := newMinion("Anub'arak", 9, 8, 4)
	anubarak.addDeathrattle(&anubarakDeathrattle{anubarak})
	return anubarak
}
