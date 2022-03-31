package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type UndoDraco struct {
}

func (p UndoDraco) Name() string {
	return "undodraco"
}

func (p UndoDraco) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.CancelItem(l, c)(2210016)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(240000110, 2)
	return true
}
