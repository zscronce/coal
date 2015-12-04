package coal

func dumpCharacters(slice []character) {
	print("[")

	for _, c := range slice[:len(slice)-1] {
		dumpCharacter(c)
		print(", ")
	}
	print("\b\b")
	println("]")
}

func dumpCharacter(c character) {
	print("{", c.name(), c.cost(), c.attack(), c.health(), c.maxHealth(), "}")
}
