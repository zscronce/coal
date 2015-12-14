package coal

//import "fmt"

type aura interface {
	register(Game)
	unRegister(Game)
}

type onEndTurn struct {
	ef effect
}

func (this *onEndTurn) register(g Game) {
	g.onEndTurn(this.ef)
}

func (this *onEndTurn) unRegister(g Game) {
	g.offEndTurn(this.ef)
}

type onDamage struct {
	ef effect
}

func (this *onDamage) register(g Game) {
	g.onDamage(this.ef)
}

func (this *onDamage) unRegister(g Game) {
	g.offDamage(this.ef)
}

type effect interface {
	apply(Game, ...interface{})
}

// General-purpose effect
type lambdaEffect struct {
	lambda func(Game, ...interface{})
}

func (this *lambdaEffect) apply(g Game, params ...interface{}) {
	this.lambda(g, params)
}

// Buffs or debuffs the attack
type addAttack struct {
	delta int
}

func (this *addAttack) getCharacter(g Game, params ...interface{}) Character {
	return g.Players()[params[0].(int)].characterAt(params[1].(int))
}

func (this *addAttack) apply(g Game, params ...interface{}) {
	this.applyTo(this.getCharacter(g, params), this.makeMutation())
}

func (this *addAttack) makeMutation() mutation {
	return &additiveMutation{this.delta}
}

func (this *addAttack) applyTo(target Character, mut mutation) {
	target.mutateAttack(mut)
}

// Buffs or debuffs the attack of a character for a single turn
type tempAddAttack struct {
	delta int
}

func (this *tempAddAttack) apply(g Game, params ...interface{}) {
	addAtk := &addAttack{this.delta}
	target := addAtk.getCharacter(g, params...)
	mut := addAtk.makeMutation()
	addAtk.applyTo(target, mut)

	removeChangeAttack := &lambdaEffect{}
	removeChangeAttack.lambda = func(_ Game, _ ...interface{}) {
		target.unMutateAttack(mut)
		g.offEndTurn(removeChangeAttack)
	}

	g.onEndTurn(removeChangeAttack)
}

// Does a flat amount of damage to all characters
type damageAllCharacters struct {
	damage int
}

func (this *damageAllCharacters) apply(g Game, params ...interface{}) {
	g.damageAllCharacters(this.damage)
}

type destroyWeapon struct {
	pIdx int
}

func (this *destroyWeapon) apply(g Game, params ...interface{}) {
	g.equip(this.pIdx, nil)
}
