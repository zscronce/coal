package coal

type delta struct {
	f func(int) int
}

type card interface {
	name() string
	cost() int
	play(*game, ...int)
	changeCost(*delta)
}

type character interface {
	card
	attack() int
	health() int
	maxHealth() int
	damage(int)
	isDead() bool
	changeAttack(*delta)
	unChangeAttack(*delta)
}

type hero interface {
	character
	armor() int
	weapon() weapon
}

type minion interface {
	character
	setBattlecry(effect)
	battlecry(*game, ...int)
	addDeathrattle(effect)
	deathrattle(*game)
	changeMaxHealth(*delta)
}

type spell interface {
	card
}

type trap interface {
	spell
}

type weapon interface {
	card
	attack() int
	durability() int
}

type cardImpl struct {
	name_      string
	cost_      int
	costDeltas []*delta
}

type characterImpl struct {
	cardImpl
	attack_      int
	health_      int
	maxHealth_   int
	attackDeltas []*delta
}

type heroImpl struct {
	characterImpl
	armor_  int
	weapon_ weapon
}

type minionImpl struct {
	characterImpl
	maxHealthDeltas []*delta
	battlecry_      effect
	deathrattle_    []effect
}

type spellImpl struct {
	cardImpl
}

type trapImpl struct {
	spellImpl
}

type weaponImpl struct {
	card
	attack_     int
	durability_ int
}

func newHero(name string) hero {
	return &heroImpl{
		characterImpl: characterImpl{
			cardImpl: cardImpl{
				name_: name,
				cost_: 2,
			},
			attack_:    0,
			health_:    30,
			maxHealth_: 30,
		},
		armor_:  0,
		weapon_: nil,
	}
}

func newMinion(name string, cost int, attack int, health int) minion {
	return &minionImpl{
		characterImpl: characterImpl{
			cardImpl: cardImpl{
				name_: name,
				cost_: cost,
			},
			attack_:    attack,
			health_:    health,
			maxHealth_: health,
		},
	}
}

func (this *cardImpl) name() string {
	return this.name_
}

func (this *cardImpl) cost() int {
	c := this.cost_

	for _, d := range this.costDeltas {
		c = d.f(c)
	}

	return c
}

func (this *cardImpl) changeCost(d *delta) {
	this.costDeltas = append(this.costDeltas, d)
}

func (this *cardImpl) play(g *game, params ...int) {
	g.players[0].mana -= this.cost()
}

func (this *characterImpl) attack() int {
	a := this.attack_

	for _, d := range this.attackDeltas {
		a = d.f(a)
	}

	return a
}

func (this *characterImpl) health() int {
	return this.health_
}

func (this *characterImpl) maxHealth() int {
	return this.maxHealth_
}

func (this *characterImpl) isDead() bool {
	return this.health_ <= 0
}

func (this *characterImpl) damage(d int) {
	this.health_ -= d
}

func (this *characterImpl) changeAttack(d *delta) {
	this.attackDeltas = append(this.attackDeltas, d)
}

func (this *characterImpl) unChangeAttack(d *delta) {
	for i, attackDelta := range this.attackDeltas {
		if d == attackDelta {
			this.attackDeltas = append(this.attackDeltas[:i], this.attackDeltas[i+1:]...)
		}
	}
}

func (this *minionImpl) play(g *game, params ...int) {
	this.cardImpl.play(g)
	g.summon(0, this)
	this.battlecry(g, params...)
}

func (this *minionImpl) setBattlecry(bc effect) {
	this.battlecry_ = bc
}

func (this *minionImpl) addDeathrattle(dr effect) {
	this.deathrattle_ = append(this.deathrattle_, dr)
}

func (this *minionImpl) battlecry(g *game, params ...int) {
	if this.battlecry_ != nil {
		this.battlecry_.apply(g, params...)
	}
}

func (this *minionImpl) deathrattle(g *game) {
	for _, dr := range this.deathrattle_ {
		dr.apply(g)
	}
}

func (this *minionImpl) changeMaxHealth(d *delta) {
	this.maxHealthDeltas = append(this.maxHealthDeltas, d)
}

func (this *heroImpl) weapon() weapon {
	return this.weapon_
}

func (this *heroImpl) armor() int {
	return this.armor_
}
