package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterEarth01 struct {
}

func (p EnterEarth01) Name() string {
	return "enter_earth01"
}

func (p EnterEarth01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.HasItem(l, c)(4031890) {
		processor.SendPinkNotice(l, c)("WARP_CARD_NEEDED")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(120000101, "earth01")
	return true
}
