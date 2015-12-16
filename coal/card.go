package coal

type Card interface {
	Name() string
	Cost() int
	play(Game, ...interface{})
}

type Character interface {
	Card
	Attack() int
	Health() int
	MaxHealth() int
	damage(int) int
	mutateAttack(mutation)
	unMutateAttack(mutation)
}

type Hero interface {
	Character
	Armor() int
	Weapon() Weapon
	equip(Weapon)
}

type Minion interface {
	Character
	Charge() bool
	DivineShield() bool
	Taunt() bool
	Windfury() bool
	addDeathrattle(effect)
	getDeathrattles() []effect
	getAuras() []aura
}

type Spell interface {
	Card
}

type Trap interface {
	Spell
}

type Weapon interface {
	Card
	Attack() int
	Durability() int
	degrade()
	getDeathrattles() []effect
}

type deathrattleCard interface {
	Card
	getDeathrattles() []effect
}

type card struct {
	name string
	cost int
}

type character struct {
	card
	attack          int
	health          int
	maxHealth       int
	attackMutations []mutation
}

type hero struct {
	character
	armor  int
	weapon Weapon
	power  effect
}

type minion struct {
	character
	charge       bool
	divineShield bool
	taunt        bool
	windfury     bool
	battlecry    effect
	deathrattle  []effect
	auras        []aura
}

type spell struct {
	card
}

type trap struct {
	spell
}

type weapon struct {
	card
	attack      int
	durability  int
	battlecry   effect
	deathrattle effect
}

func newHero(name string) *hero {
	return &hero{
		character: character{
			card: card{
				name: name,
				cost: 2,
			},
			attack:    0,
			health:    30,
			maxHealth: 30,
		},
		armor:  0,
		weapon: nil,
	}
}

func newMinion(name string, cost int, attack int, health int) *minion {
	return &minion{
		character: character{
			card: card{
				name: name,
				cost: cost,
			},
			attack:    attack,
			health:    health,
			maxHealth: health,
		},
		charge:       false,
		divineShield: false,
		taunt:        false,
		windfury:     false,
	}
}

func newWeapon(name string, cost int, attack int, durability int) *weapon {
	return &weapon{
		card: card{
			name: name,
			cost: cost,
		},
		attack:     attack,
		durability: durability,
	}
}

func (this *card) Name() string {
	return this.name
}

func (this *card) Cost() int {
	return this.cost
}

func (this *character) Attack() int {
	a := this.attack
	for _, mut := range this.attackMutations {
		a = mut.apply(a)
	}
	return a
}

func (this *character) Health() int {
	return this.health
}

func (this *character) MaxHealth() int {
	return this.maxHealth
}

func (this *character) damage(dmg int) int {
	this.health -= dmg
	return dmg
}

func (this *character) mutateAttack(mut mutation) {
	this.attackMutations = append(this.attackMutations, mut)
}

func (this *character) unMutateAttack(mut mutation) {
	for i, m := range this.attackMutations {
		if m == mut {
			this.attackMutations = append(this.attackMutations[:i], this.attackMutations[i+1:]...)
			return
		}
	}
}

func (this *hero) Weapon() Weapon {
	return this.weapon
}

func (this *hero) Armor() int {
	return this.armor
}

func (this *hero) damage(dmg int) int {
	dmgToArmor := dmg

	if dmg > this.Armor() {
		dmgToArmor = this.Armor()
	}

	this.armor -= dmgToArmor
	dmgToHealth := dmg - dmgToArmor
	this.character.damage(dmgToHealth)

	return dmg
}

func (this *hero) equip(w Weapon) {
	this.weapon = w
}

func (this *hero) play(g Game, params ...interface{}) {
	this.power.apply(g, params...)
}

func (this *minion) Charge() bool {
	return this.charge
}

func (this *minion) DivineShield() bool {
	return this.divineShield
}

func (this *minion) Taunt() bool {
	return this.taunt
}

func (this *minion) Windfury() bool {
	return this.windfury
}

func (this *minion) addAura(a aura) {
	this.auras = append(this.auras, a)
}

func (this *minion) addDeathrattle(dr effect) {
	this.deathrattle = append(this.deathrattle, dr)
}

func (this *minion) damage(dmg int) int {
	if dmg == 0 {
		return dmg
	}

	if this.DivineShield() {
		this.divineShield = false
		return 0
	}

	return this.character.damage(dmg)
}

func (this *minion) getDeathrattles() []effect {
	if this.deathrattle == nil {
		return []effect{}
	} else {
		return this.deathrattle
	}
}

func (this *minion) getAuras() []aura {
	return this.auras
}

func (this *minion) play(g Game, params ...interface{}) {
	g.summon(0, this)
	if this.battlecry != nil {
		this.battlecry.apply(g, params...)
	}
}

func (this *weapon) Attack() int {
	return this.attack
}

func (this *weapon) Durability() int {
	return this.durability
}

func (this *weapon) degrade() {
	this.durability--
}

func (this *weapon) getDeathrattles() []effect {
	if this.deathrattle != nil {
		return []effect{this.deathrattle}
	} else {
		return []effect{}
	}
}

func (this *weapon) play(g Game, params ...interface{}) {
	g.equip(0, this)
	if this.battlecry != nil {
		this.battlecry.apply(g, params)
	}
}
