package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MasteriaCC6A struct {
}

func (p MasteriaCC6A) Name() string {
	return "Masteria_CC6_A"
}

func (p MasteriaCC6A) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(610020010, "CC1_A")
	return true
}
