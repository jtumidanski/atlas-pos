package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal03 struct {
}

func (p GlpqPortal03) Name() string {
	return "glpqPortal03"
}

func (p GlpqPortal03) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.MeetsCriteria(l, span)(c.CharacterId(), character.IsJobNiche(4)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(610030530, 0)
		return true
	}
	script.SendPinkNotice(l, c)("THIEF_ONLY")
	return false
}
