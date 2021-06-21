package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type CaptinsG00 struct {
}

func (p CaptinsG00) Name() string {
	return "captinsg00"
}

func (p CaptinsG00) Enter(l logrus.FieldLogger, c script.Context) bool {
	if !character.HasItem(l, c)(4000381) {
		character.SendPinkNotice(l, c)("CAPTAIN_LATANICA_MISSING_ESSENCE")
		return false
	}

	party, ok := character.GetParty(l, c)
	if !ok {
		character.SendPinkNotice(l, c)("BOSS_PARTY_NEEDED")
		return false
	}

	if !character.PartyLeader(l, c) {
		character.SendPinkNotice(l, c)("BOSS_PARTY_LEADER_START")
		return false
	}

	err := script.StartEvent(l, c)("LatanicaBattle", party.Id(), c.MapId(), 1)
	if err != nil {
		character.SendPinkNotice(l, c)("BOSS_ALREADY_STARTED")
		return false
	}

	character.PlayPortalSound(l)
	return true
}