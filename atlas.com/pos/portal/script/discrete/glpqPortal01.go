package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal01 struct {
}

func (p GlpqPortal01) Name() string {
	return "glpqPortal01"
}

func (p GlpqPortal01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.MeetsCriteria(l, span)(c.CharacterId(), character.IsJobNiche(3)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(610030540, 0)
		return true
	}
	script.SendPinkNotice(l, c)("ARCHER_ONLY")
	return false
}
