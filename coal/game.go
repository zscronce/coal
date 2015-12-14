package coal

// game interface
type Game interface {
	Active() *Player
	Inactive() *Player
	Players() [2]*Player
	PlayFromHand(int, ...interface{})
	PlayHeroPower(...interface{})
	Attack(int, int)
	EndTurn()
	addToHand(int, Card)
	damageAllCharacters(int)
	damageHero(int, int)
	destroy(Minion)
	destroyAt(int, int)
	draw(int, int)
	drawOne(int)
	equip(int, Weapon)
	ownerOf(Card) int
	summon(int, Minion)
	onEndTurn(effect)
	offEndTurn(effect)
	onDamage(effect)
	offDamage(effect)
}

// simple container to hold player state
type Player struct {
	hero    Hero
	minions []Minion
	traps   []Trap
	hand    []Card
	deck    []Card
	mana    int
	maxMana int
	fatigue int
}

type game struct {
	players        [2]*Player
	runDepth       int
	owner          map[Card]int
	deadIdx        [2][]bool
	endTurnEffects []effect
	damageEffects  []effect
}

func NewGame(players [2]*Player) Game {
	g := &game{
		players:        players,
		runDepth:       0,
		owner:          map[Card]int{},
		deadIdx:        [2][]bool{[]bool{}, []bool{}},
		endTurnEffects: []effect{},
		damageEffects:  []effect{},
	}

	for p, pl := range g.Players() {
		g.owner[pl.hero] = p
		for m := range pl.minions {
			g.intakeMinion(p, m)
		}
	}

	return g
}

func (this *Player) characterAt(idx int) Character {
	if idx == 0 {
		return this.hero
	} else {
		return this.minions[minionIdx(idx)]
	}
}

func (this *game) Active() *Player {
	return this.Players()[0]
}

func (this *game) Inactive() *Player {
	return this.Players()[1]
}

func (this *game) Players() [2]*Player {
	return this.players
}

// Plays the card at this.Players()[0].hand[h], optional additional parameters (target selection)
func (this *game) PlayFromHand(h int, params ...interface{}) {
	cd := this.Active().hand[h]
	before := this.Active().hand[:h]
	after := this.Active().hand[h+1:]
	this.Active().hand = append(before, after...)
	this.play(cd, params...)
}

func (this *game) PlayHeroPower(params ...interface{}) {
	this.play(this.Active().hero, params...)
}

func (this *game) EndTurn() {
	for e := 0; e < len(this.endTurnEffects); e++ {
		this.endTurnEffects[e].apply(this)
	}

	this.players[0], this.players[1] = this.players[1], this.players[0]
	this.deadIdx[0], this.deadIdx[1] = this.deadIdx[1], this.deadIdx[0]
	for c, o := range this.owner {
		this.owner[c] = o ^ 1
	}
}

func (this *game) Attack(aIdx int, dIdx int) {
	attacker := this.Active().characterAt(aIdx)
	defender := this.Inactive().characterAt(dIdx)

	var attackerWeap Weapon = nil
	if attacker.(interface{}) == this.Active().hero.(interface{}) {
		attackerWeap = this.Active().hero.Weapon()
	}

	attackerAttack := attacker.Attack()
	defenderAttack := defender.Attack()

	if attackerWeap != nil {
		attackerAttack += attackerWeap.Attack()
	}

	if defender.(interface{}) == this.Inactive().hero.(interface{}) {
		defenderAttack = 0
	}

	damagesPlayerIdx := []int{1, 0}
	damagesCharacterIdx := []int{dIdx, aIdx}
	damagesAmount := []int{attackerAttack, defenderAttack}

	this.damage(damagesPlayerIdx, damagesCharacterIdx, damagesAmount)

	if attackerWeap != nil {
		attackerWeap.degrade()
		if attackerWeap.Durability() == 0 {
			this.equip(0, nil)
		}
	}
}

func (this *game) addToHand(p int, cd Card) {
	this.Players()[p].hand = append(this.Players()[p].hand, cd)
	this.owner[cd] = p
}

func (this *game) damageAllCharacters(dmg int) {
	pIdx := []int{}
	cIdx := []int{}
	amount := []int{}

	for p, pl := range this.Players() {
		pIdx = append(pIdx, p)
		cIdx = append(cIdx, 0)
		amount = append(amount, dmg)

		for m := range pl.minions {
			pIdx = append(pIdx, p)
			cIdx = append(cIdx, characterIdx(m))
			amount = append(amount, dmg)
		}
	}

	this.damage(pIdx, cIdx, amount)
}

func (this *game) damageHero(p int, dmg int) {
	this.damage([]int{p}, []int{0}, []int{dmg})
}

func (this *game) destroy(target Minion) {
	for p, pl := range this.Players() {
		for m, min := range pl.minions {
			if min == target {
				this.destroyAt(p, m)
			}
		}
	}
}

func (this *game) destroyAt(p int, m int) {
	this.run(func() {
		this.deadIdx[p][m] = true
	})
}

func (this *game) draw(p int, n int) {
	for i := 0; i < n; i++ {
		this.drawOne(p)
	}
}

func (this *game) drawOne(p int) {
	this.run(func() {
		pl := this.Players()[p]

		if len(pl.deck) == 0 {
			pl.fatigue++
			this.damageHero(p, pl.fatigue)
		} else {
			drawn := pl.deck[0]
			pl.deck = pl.deck[1:]
			pl.hand = append(pl.hand, drawn)
		}
	})
}

func (this *game) equip(p int, wp Weapon) {
	prev := this.Players()[p].hero.Weapon()
	this.Players()[p].hero.equip(wp)
	if prev != nil {
		this.deathrattle(prev)
	}
}

func (this *game) play(cd Card, params ...interface{}) {
	this.Active().mana -= cd.Cost()
	cd.play(this, params...)
}

func (this *game) ownerOf(cd Card) int {
	owner, hasOwner := this.owner[cd]

	if !hasOwner {
		return -1
	}

	return owner
}

func (this *game) summon(p int, min Minion) {
	pl := this.Players()[p]
	pl.minions = append(pl.minions, min)
	this.intakeMinion(p, len(pl.minions)-1)
}

func (this *game) onEndTurn(e effect) {
	addEffect(&this.endTurnEffects, e)
}

func (this *game) offEndTurn(e effect) {
	removeEffect(&this.endTurnEffects, e)
}

func (this *game) onDamage(e effect) {
	addEffect(&this.damageEffects, e)
}

func (this *game) offDamage(e effect) {
	removeEffect(&this.damageEffects, e)
}

func (this *game) damage(pIdx []int, cIdx []int, amounts []int) {
	this.run(func() {
		tookDamage := []Character{}
		tookAmount := []int{}

		for i := range amounts {
			if amounts[i] == 0 {
				continue
			}

			mIdx := minionIdx(cIdx[i])
			if mIdx >= 0 && this.deadIdx[pIdx[i]][mIdx] {
				continue
			}

			ch := this.Players()[pIdx[i]].characterAt(cIdx[i])
			taken := ch.damage(amounts[i])
			if taken != 0 {
				tookDamage = append(tookDamage, ch)
				tookAmount = append(tookAmount, taken)
			}

			if mIdx >= 0 && ch.Health() <= 0 {
				this.deadIdx[pIdx[i]][mIdx] = true
			}
		}

		for i := range tookDamage {
			for _, e := range this.damageEffects {
				e.apply(this, tookDamage[i], tookAmount[i])
			}
		}
	})
}

func (this *game) deathrattle(cd deathrattleCard) {
	for _, e := range cd.getDeathrattles() {
		this.run(func() {
			e.apply(this)
		})
	}
}

// Should be called whenever a new minion is added to a player's minion line. Performs upkeep of the
// game's state.
func (this *game) intakeMinion(p int, m int) {
	min := this.Players()[p].minions[m]

	this.owner[min] = p
	this.deadIdx[p] = append(append(this.deadIdx[p][:m], false), this.deadIdx[p][m:]...)
	for _, aura := range min.getAuras() {
		aura.register(this)
	}
}

// Should be called whenever a minion is removed from a player's minion line. Performs upkeep of the
// game's state.
func (this *game) outturnMinion(p int, m int) {
	this.deadIdx[p] = append(this.deadIdx[p][:m], this.deadIdx[p][m+1:]...)
	for _, aura := range this.Players()[p].minions[m].getAuras() {
		aura.unRegister(this)
	}
}

func (this *game) run(phase func()) {
	this.runDepth++
	phase()
	this.runDepth--

	if this.runDepth != 0 {
		return
	}

	for p := range this.Players() {
		for m := 0; m < len(this.Players()[p].minions); m++ {
			if this.deadIdx[p][m] {
				min := this.Players()[p].minions[m]
				this.outturnMinion(p, m)
				this.Players()[p].minions = append(this.Players()[p].minions[:m], this.Players()[p].minions[m+1:]...)
				this.deathrattle(min)
				m--
			}
		}
	}
}

func addEffect(slice *[]effect, e effect) {
	*slice = append(*slice, e)
}

func characterIdx(minionIdx int) int {
	return minionIdx + 1
}

func minionIdx(characterIdx int) int {
	return characterIdx - 1
}

func removeEffect(slice *[]effect, e effect) {
	for i := 0; i < len(*slice); i++ {
		if (*slice)[i] == e {
			*slice = append((*slice)[:i], (*slice)[i+1:]...)
		}
	}
}
