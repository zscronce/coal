package coal

type player struct {
	hero    hero
	minions []minion
	traps   []trap
	hand    []card
	deck    []card
	mana    int
	maxMana int
}

type game struct {
	players       [2]*player
	runDepth      int
	owner         map[card]int
	onceNextTurn_ []func()
	onNextTurn_   []func()
}

type damage struct {
	playerIdx    int
	characterIdx int
	amount       int
}

func newGame(players [2]*player) *game {
	g := &game{
		players: players,
		owner:   map[card]int{},
	}

	for pIdx, p := range g.players {
		for _, c := range p.characters() {
			g.owner[c] = pIdx
		}
	}

	return g
}

func (this *player) characters() []character {
	characters := []character{}

	characters = append(characters, this.hero)
	for _, minion := range this.minions {
		characters = append(characters, minion)
	}

	return characters
}

func (this *game) characters() [2][]character {
	return [2][]character{
		this.players[0].characters(),
		this.players[1].characters(),
	}
}

func (this *game) play(h int, params ...int) {
	c := this.players[0].hand[h]
	this.players[0].hand = append(this.players[0].hand[:h], this.players[0].hand[h+1:]...)
	c.play(this, params...)
}

func (this *game) nextTurn() {
	runCallbacks(this.onNextTurn_, &this.onceNextTurn_)

	this.players[0], this.players[1] = this.players[1], this.players[0]
	for c, o := range this.owner {
		this.owner[c] = o ^ 1
	}
}

func (this *game) attack(aIdx int, iIdx int) {
	characters := this.characters()
	attacker := characters[0][aIdx]
	defender := characters[1][iIdx]

	damages := []damage{
		damage{1, iIdx, attacker.attack()},
		damage{0, aIdx, defender.attack()},
	}

	this.damage(damages)
}

func (this *game) damage(damages []damage) {
	this.run(func() {
		characters := this.characters()

		for _, d := range damages {
			characters[d.playerIdx][d.characterIdx].damage(d.amount)
		}
	})
}

func (this *game) summon(p int, m minion) {
	this.players[p].minions = append(this.players[p].minions, m)
	this.owner[m] = p
}

func (this *game) addToHand(p int, c card) {
	this.players[p].hand = append(this.players[p].hand, c)
	this.owner[c] = p
}

func (this *game) run(phase func()) {
	this.runDepth++
	phase()
	this.runDepth--

	if this.runDepth == 0 {
		dead := []minion{}

		// aggregate list of dead minions, and remove them from players' minion lines
		for _, p := range this.players {
			for m := 0; m < len(p.minions); m++ {
				if p.minions[m].isDead() {
					dead = append(dead, p.minions[m])
					p.minions = append(p.minions[:m], p.minions[m+1:]...)
					m--
				}
			}
		}

		for _, m := range dead {
			m.deathrattle(this)
		}
	}
}

func runCallbacks(callbacks []func(), onceCallbacks *[]func()) {
	for _, callback := range append(callbacks, *onceCallbacks...) {
		callback()
	}
	*onceCallbacks = []func(){}
}
