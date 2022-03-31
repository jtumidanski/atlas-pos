package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Male00 struct {
}

func (p Male00) Name() string {
	return "male00"
}

func (p Male00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.GetGender(l, span)(c.CharacterId()) == character.GenderMale {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(670010200, 3)
		return true
	}
	processor.SendPinkNotice(l, c)("CANNOT_PROCEED")
	return false
}
