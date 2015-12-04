package coal

func newAbomination() minion {
	abomination := newMinion("Abomination", 5, 4, 4)
	abomination.addDeathrattle(&damageAllCharacters{2})
	return abomination
}
