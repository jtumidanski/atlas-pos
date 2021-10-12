package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type BalogEnd struct {
}

func (p BalogEnd) Name() string {
	return "balog_end"
}

func (p BalogEnd) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.CanHold(l, c)(4001261, 1) {
		script.GainItem(l, c)(4001261, 1)
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(105100100, 0)
		return true
	}
	script.SendPinkNotice(l, c)("INVENTORY_FULL_ERROR")
	return false
}
