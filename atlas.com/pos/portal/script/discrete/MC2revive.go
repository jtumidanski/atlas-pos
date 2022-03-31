package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MC2Revive struct {
}

func (p MC2Revive) Name() string {
	return "MC2revive"
}

func (p MC2Revive) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpRandom(l, span, c)(c.MapId() - 100)
	return true
}
