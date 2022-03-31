package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCC1A struct {
}

func (p MasteriaCC1A) Name() string {
	return "Masteria_CC1_A"
}

func (p MasteriaCC1A) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(610020015, "CC6_A")
	return true
}
