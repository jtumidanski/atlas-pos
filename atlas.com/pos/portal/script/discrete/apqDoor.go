package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ApqDoor struct {
}

func (p ApqDoor) Name() string {
	return "apqDoor"
}

func (p ApqDoor) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	r, err := processor.GetReactor(l, c)(c.MapId(), "gate"+c.Portal().Name())
	if err != nil || r.State() != 4 {
		processor.SendPinkNotice(l, c)("GATE_NOT_YET_OPEN")
		return false
	}

	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(670010600, "gt"+c.Portal().Name()+"PIB")
	return true
}
