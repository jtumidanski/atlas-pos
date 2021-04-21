package script

import (
	"github.com/sirupsen/logrus"
)

type captinsG00 struct {
}

func CaptinsG00() PortalScript {
	return captinsG00{}
}

func (a captinsG00) Name() string {
	return "captinsg00"
}

func (a captinsG00) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if !p.HasItem(4000381) {
		p.SendPinkNotice("CAPTAIN_LATANICA_MISSING_ESSENCE")
		return false
	}

	party, ok := p.GetParty()
	if !ok {
		p.SendPinkNotice("BOSS_PARTY_NEEDED")
		return false
	}

	if !p.PartyLeader() {
		p.SendPinkNotice("BOSS_PARTY_LEADER_START")
		return false
	}

	err := p.StartEvent("LatanicaBattle", party.Id(), context.MapId(), 1)
	if err != nil {
		p.SendPinkNotice("BOSS_ALREADY_STARTED")
		return false
	}

	p.PlayPortalSound()
	return true
}
