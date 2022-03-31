package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type BalogEnd struct {
}

func (p BalogEnd) Name() string {
	return "balog_end"
}

func (p BalogEnd) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.CanHold(l, c)(4001261, 1) {
		processor.GainItem(l, c)(4001261, 1)
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(105100100, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("INVENTORY_FULL_ERROR")
	return false
}
