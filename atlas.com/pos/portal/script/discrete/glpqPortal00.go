package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GlpqPortal00 struct {
}

func (p GlpqPortal00) Name() string {
	return "glpqPortal00"
}

func (p GlpqPortal00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.MeetsCriteria(l, span)(c.CharacterId(), character.IsJobNiche(1)) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(610030510, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("WARRIOR_ONLY")
	return false
}
