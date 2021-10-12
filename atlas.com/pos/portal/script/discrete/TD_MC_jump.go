package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type TDMCJump struct {
}

func (p TDMCJump) Name() string {
	return "TD_MC_jump"
}

func (p TDMCJump) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(106020501, 0)
	return true
}
