package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal04 struct {
}

func (p GlpqPortal04) Name() string {
	return "glpqPortal04"
}

func (p GlpqPortal04) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.MeetsCriteria(l)(c.CharacterId(), character.IsJobNiche(5)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(610030550, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PIRATE_ONLY")
	return false
}
