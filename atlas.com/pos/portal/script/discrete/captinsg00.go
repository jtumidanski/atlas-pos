package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CaptinsG00 struct {
}

func (p CaptinsG00) Name() string {
	return "captinsg00"
}

func (p CaptinsG00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.HasItem(l, c)(4000381) {
		script.SendPinkNotice(l, c)("CAPTAIN_LATANICA_MISSING_ESSENCE")
		return false
	}

	party, ok := script.GetParty(l, c)
	if !ok {
		script.SendPinkNotice(l, c)("BOSS_PARTY_NEEDED")
		return false
	}

	if !script.PartyLeader(l, c) {
		script.SendPinkNotice(l, c)("BOSS_PARTY_LEADER_START")
		return false
	}

	err := script.StartEvent(l, c)("LatanicaBattle", party.Id(), c.MapId(), 1)
	if err != nil {
		script.SendPinkNotice(l, c)("BOSS_ALREADY_STARTED")
		return false
	}

	script.PlayPortalSound(l, c)
	return true
}
