package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Under30Gate struct {
}

func (p Under30Gate) Name() string {
	return "under30gate"
}

func (p Under30Gate) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.AboveLevel(l, span)(c.CharacterId(), 30) {
		processor.SendPinkNotice(l, c)("CANNOT_PROCEED_PAST_POINT")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(990000640, 1)
	return true
}
