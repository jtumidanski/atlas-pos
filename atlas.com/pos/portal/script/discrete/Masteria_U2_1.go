package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaU21 struct {
}

func (p MasteriaU21) Name() string {
	return "Masteria_U2_1"
}

func (p MasteriaU21) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992040) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610010004, "U5_1")
		return false
	}
	return true
}
