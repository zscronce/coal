package coal

type anubarakDeathrattle struct {
	me Minion
}

func (this *anubarakDeathrattle) apply(g Game, params ...interface{}) {
	owner := g.ownerOf(this.me)
	g.addToHand(owner, this.me)
	g.summon(owner, newMinion("Nerubian", 2, 4, 4))
}

func newAnubarak() Minion {
	anubarak := newMinion("Anub'arak", 9, 8, 4)
	anubarak.addDeathrattle(&anubarakDeathrattle{anubarak})
	return anubarak
}
