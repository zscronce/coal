package coal

func newAlAkirTheWindlord() Minion {
	alAkirTheWindlord := newMinion("Al'Akir the Windlord", 8, 3, 5)
	alAkirTheWindlord.charge = true
	alAkirTheWindlord.divineShield = true
	alAkirTheWindlord.taunt = true
	alAkirTheWindlord.windfury = true
	return alAkirTheWindlord
}
