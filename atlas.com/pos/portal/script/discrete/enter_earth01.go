package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterEarth01 struct {
}

func (p EnterEarth01) Name() string {
	return "enter_earth01"
}

func (p EnterEarth01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.HasItem(l, c)(4031890) {
		script.SendPinkNotice(l, c)("WARP_CARD_NEEDED")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(120000101, "earth01")
	return true
}
