package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GlpqPortal03 struct {
}

func (p GlpqPortal03) Name() string {
	return "glpqPortal03"
}

func (p GlpqPortal03) Enter(l logrus.FieldLogger, c script.Context) bool {
	if character.MeetsCriteria(l)(c.CharacterId(), character.IsJobNiche(4)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, c)(610030530, 0)
		return true
	}
	script.SendPinkNotice(l, c)("THIEF_ONLY")
	return false
}
