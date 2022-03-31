package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutRider struct {
}

func (p OutRider) Name() string {
	return "outRider"
}

func (p OutRider) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.CanHold(l, c)(4001193, 1) {
		processor.SendPinkNotice(l, c)("FREE_SPACE_FOR_COUSE_CLEAR_TOKEN")
		return false
	}
	processor.GainItem(l, c)(4001193, 1)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(211050000, 4)
	return true
}
