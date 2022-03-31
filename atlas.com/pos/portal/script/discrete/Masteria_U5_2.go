package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaU52 struct {
}

func (p MasteriaU52) Name() string {
	return "Masteria_U5_2"
}

func (p MasteriaU52) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(3992040) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(610010004, "U5_1")
		return false
	}
	return true
}
