package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaU21 struct {
}

func (p MasteriaU21) Name() string {
	return "Masteria_U2_1"
}

func (p MasteriaU21) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(3992040) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(610010004, "U5_1")
		return false
	}
	return true
}
