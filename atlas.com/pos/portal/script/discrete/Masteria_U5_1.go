package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaU51 struct {
}

func (p MasteriaU51) Name() string {
	return "Masteria_U5_1"
}

func (p MasteriaU51) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610010001, "sU2_1")
		return false
	}
	return true
}
