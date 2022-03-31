package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartGoFoward0 struct {
}

func (p DepartGoFoward0) Name() string {
	return "DepartFoward0"
}

func (p DepartGoFoward0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.MapId() == 103040410 && processor.QuestCompleted(l, c)(2287) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(103040420, "right00")
		return true
	}
	if c.MapId() == 103040420 && processor.QuestCompleted(l, c)(2288) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(103040430, "right00")
		return true
	}
	if c.MapId() == 103040410 && processor.QuestStarted(l, c)(2287) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(103040420, "right00")
		return true
	}
	if c.MapId() == 103040420 && processor.QuestStarted(l, c)(2288) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(103040430, "right00")
		return true
	}
	if c.MapId() == 103040440 || c.MapId() == 103040450 {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(c.MapId()+10, "right00")
		return true
	}
	processor.SendPinkNotice(l, c)("CANNOT_ACCESS")
	return false
}
