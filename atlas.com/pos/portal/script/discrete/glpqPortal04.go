package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal04 struct {
}

func (p GlpqPortal04) Name() string {
	return "glpqPortal04"
}

func (p GlpqPortal04) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.MeetsCriteria(l, span)(c.CharacterId(), character.IsJobNiche(5)) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(610030550, 0)
		return true
	}
	script.SendPinkNotice(l, c)("PIRATE_ONLY")
	return false
}
