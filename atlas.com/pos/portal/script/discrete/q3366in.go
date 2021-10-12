package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3366In struct {
}

func (p Q3366In) Name() string {
	return "q3366in"
}

func (p Q3366In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestStarted(l, c)(3366) {
		script.SendPinkNotice(l, c)("ROOM_NO_ACCESS")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(926130101, 0)
	return true
}
