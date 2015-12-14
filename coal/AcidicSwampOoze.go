package coal

func newAcidicSwampOoze() Minion {
	acidicSwampOoze := newMinion("Acidic Swamp Ooze", 2, 3, 2)
	acidicSwampOoze.battlecry = &destroyWeapon{1}
	return acidicSwampOoze
}
