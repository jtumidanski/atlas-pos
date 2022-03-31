package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MoveElin struct {
}

func (p MoveElin) Name() string {
	return "move_elin"
}

func (p MoveElin) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(300000100, "out00")
	processor.SendPinkNotice(l, c)("PASSING_TIME_GATE")
	return true
}
