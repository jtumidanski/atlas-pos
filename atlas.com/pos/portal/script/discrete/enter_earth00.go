package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterEarth00 struct {
}

func (p EnterEarth00) Name() string {
	return "enter_earth00"
}

func (p EnterEarth00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.HasItem(l, c)(4031890) {
		processor.SendPinkNotice(l, c)("WARP_CARD_NEEDED")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(221000300, "earth00")
	return true
}
