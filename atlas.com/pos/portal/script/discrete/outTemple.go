package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutTemple struct {
}

func (p OutTemple) Name() string {
	return "outTemple"
}

func (p OutTemple) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.UseItem(l, c)(2210016)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(200090510, 0)
	return true
}
