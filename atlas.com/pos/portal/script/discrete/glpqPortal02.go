package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal02 struct {
}

func (p GlpqPortal02) Name() string {
	return "glpqPortal02"
}

func (p GlpqPortal02) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.MeetsCriteria(l, span)(c.CharacterId(), character.IsJobNiche(2)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(610030521, 0)
		return true
	}
	script.SendPinkNotice(l, c)("MAGICIAN_ONLY")
	return false
}
