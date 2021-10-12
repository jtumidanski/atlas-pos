package generic

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func EnterMarket(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.MapId() != 910000000 {
		script.SaveLocation(l, c)("FREE_MARKET")
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(910000000, "out00")
		return true
	}
	return false
}
