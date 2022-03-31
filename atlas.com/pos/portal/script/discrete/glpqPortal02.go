package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
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
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(610030521, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("MAGICIAN_ONLY")
	return false
}
