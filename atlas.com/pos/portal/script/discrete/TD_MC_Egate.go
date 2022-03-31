package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCEGate struct {
}

func (p TDMCEGate) Name() string {
	return "TD_MC_Egate"
}

func (p TDMCEGate) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(106021300, 1)
	return true
}
