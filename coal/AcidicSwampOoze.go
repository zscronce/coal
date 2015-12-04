package coal

type acidicSwampOozeBattlecry struct {
}

func (this *acidicSwampOozeBattlecry) apply(g *game, params ...int) {
	g.equip(1, nil)
}

func newAcidicSwampOoze() minion {
	acidicSwampOoze := newMinion("Acidic Swamp Ooze", 2, 3, 2)
	acidicSwampOoze.battlecry_ = &acidicSwampOozeBattlecry{}
	return acidicSwampOoze
}
