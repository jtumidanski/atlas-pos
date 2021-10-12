package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type UndoDraco struct {
}

func (p UndoDraco) Name() string {
	return "undodraco"
}

func (p UndoDraco) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.CancelItem(l, c)(2210016)
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(240000110, 2)
	return true
}
