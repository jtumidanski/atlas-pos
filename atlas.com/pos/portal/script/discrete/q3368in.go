package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3368In struct {
}

func (p Q3368In) Name() string {
	return "q3368in"
}

func (p Q3368In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestStarted(l, c)(3368) {
		script.SendPinkNotice(l, c)("ROOM_NO_ACCESS")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(926130103, 0)
	return true
}
