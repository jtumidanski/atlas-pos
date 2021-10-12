package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MoveElin struct {
}

func (p MoveElin) Name() string {
	return "move_elin"
}

func (p MoveElin) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(300000100, "out00")
	script.SendPinkNotice(l, c)("PASSING_TIME_GATE")
	return true
}
