package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCEGate struct {
}

func (p TDMCEGate) Name() string {
	return "TD_MC_Egate"
}

func (p TDMCEGate) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(106021300, 1)
	return true
}
