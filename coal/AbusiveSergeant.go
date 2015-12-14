package coal

func newAbusiveSergeant() Minion {
	abusiveSergeant := newMinion("Abusive Sergeant", 1, 2, 1)
	abusiveSergeant.battlecry = &tempAddAttack{2}
	return abusiveSergeant
}
