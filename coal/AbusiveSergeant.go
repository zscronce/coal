package coal

func newAbusiveSergeant() minion {
	abusiveSergeant := newMinion("Abusive Sergeant", 1, 2, 1)
	abusiveSergeant.battlecry_ = newTempAttackDelta(2)
	return abusiveSergeant
}
