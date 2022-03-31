package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Q3368In struct {
}

func (p Q3368In) Name() string {
	return "q3368in"
}

func (p Q3368In) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(3368) {
		processor.SendPinkNotice(l, c)("ROOM_NO_ACCESS")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(926130103, 0)
	return true
}
