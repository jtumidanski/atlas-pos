package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterEarth00 struct {
}

func (p EnterEarth00) Name() string {
	return "enter_earth00"
}

func (p EnterEarth00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.HasItem(l, c)(4031890) {
		script.SendPinkNotice(l, c)("WARP_CARD_NEEDED")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(221000300, "earth00")
	return true
}
