package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TempleEnter struct {
}

func (p TempleEnter) Name() string {
	return "templeenter"
}

func (p TempleEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.CancelItem(l, c)(2210016)
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(270000100, "out00")
	return true
}
