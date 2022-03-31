package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaB31 struct {
}

func (p MasteriaB31) Name() string {
	return "Masteria_B3_1"
}

func (p MasteriaB31) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992040) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610010005, "sU6_1")
		return true
	}
	return false
}
