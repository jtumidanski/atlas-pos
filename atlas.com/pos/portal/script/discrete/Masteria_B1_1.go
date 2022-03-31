package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaB11 struct {
}

func (p MasteriaB11) Name() string {
	return "Masteria_B1_1"
}

func (p MasteriaB11) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992040) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610010005, "sU6_1")
		return true
	}
	return false
}
