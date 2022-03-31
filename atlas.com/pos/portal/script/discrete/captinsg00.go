package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CaptinsG00 struct {
}

func (p CaptinsG00) Name() string {
	return "captinsg00"
}

func (p CaptinsG00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.HasItem(l, c)(4000381) {
		processor.SendPinkNotice(l, c)("CAPTAIN_LATANICA_MISSING_ESSENCE")
		return false
	}

	party, ok := processor.GetParty(l, c)
	if !ok {
		processor.SendPinkNotice(l, c)("BOSS_PARTY_NEEDED")
		return false
	}

	if !processor.PartyLeader(l, c) {
		processor.SendPinkNotice(l, c)("BOSS_PARTY_LEADER_START")
		return false
	}

	err := processor.StartEvent(l, c)("LatanicaBattle", party.Id(), c.MapId(), 1)
	if err != nil {
		processor.SendPinkNotice(l, c)("BOSS_ALREADY_STARTED")
		return false
	}

	processor.PlayPortalSound(l, c)
	return true
}
