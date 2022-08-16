package generic

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func EnterMarket(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.MapId() != 910000000 {
		processor.SaveLocation(l, span, c)("FREE_MARKET")
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(910000000, "out00")
		return true
	}
	return false
}
