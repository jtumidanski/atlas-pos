package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaB31 struct {
}

func (p MasteriaB31) Name() string {
	return "Masteria_B3_1"
}

func (p MasteriaB31) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610010005, "sU6_1")
		return true
	}
	return false
}
