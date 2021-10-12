package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaU31 struct {
}

func (p MasteriaU31) Name() string {
	return "Masteria_U3_1"
}

func (p MasteriaU31) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610010201, "sB2_1")
		return false
	}
	return true
}
