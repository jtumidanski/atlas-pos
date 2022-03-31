package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
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
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(610030550, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PIRATE_ONLY")
	return false
}
