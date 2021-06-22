package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal00 struct {
}

func (p GlpqPortal00) Name() string {
	return "glpqPortal00"
}

func (p GlpqPortal00) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.MeetsCriteria(l)(c.CharacterId(), character.IsJobNiche(1)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(610030510, 0)
		return true
	}
	script.SendPinkNotice(l, c)("WARRIOR_ONLY")
	return false
}
