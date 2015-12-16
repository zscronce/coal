package coal

import (
	"math/rand"
)

type alarmOBotEffect struct {
	me Minion
}

func (this *alarmOBotEffect) apply(g Game, params ...interface{}) {
	pIdx := g.ownerOf(this.me)

	if pIdx != 0 {
		return
	}

	pl := g.Players()[pIdx]
	minionsInHandIdx := []int{}
	minionsInHand := []Minion{}

	for c, cd := range pl.hand {
		if min, isMinion := cd.(Minion); isMinion {
			minionsInHandIdx = append(minionsInHandIdx, c)
			minionsInHand = append(minionsInHand, min)
		}
	}

	if len(minionsInHand) == 0 {
		return
	}

	meIdx := -1
	for m, min := range pl.minions {
		if min == this.me {
			meIdx = m
		}
	}

	r := rand.Int() % len(minionsInHand)
	randMinionInHandIdx := minionsInHandIdx[r]
	randMinionInHand := minionsInHand[r]

	g.returnMinionToHand(pIdx, meIdx)
	g.removeFromHand(pIdx, randMinionInHandIdx)
	g.summon(pIdx, randMinionInHand)
}

func newAlarmOBot() Minion {
	alarmOBot := newMinion("Alarm-o-Bot", 3, 0, 3)
	alarmOBot.auras = []aura{&onStartTurn{&alarmOBotEffect{alarmOBot}}}
	return alarmOBot
}
