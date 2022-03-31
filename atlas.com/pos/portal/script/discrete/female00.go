package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Female00 struct {
}

func (p Female00) Name() string {
	return "female00"
}

func (p Female00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.GetGender(l, span)(c.CharacterId()) == character.GenderFemale {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(670010200, 4)
		return true
	}
	processor.SendPinkNotice(l, c)("CANNOT_PROCEED")
	return false
}
