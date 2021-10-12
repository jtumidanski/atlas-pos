package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DepartGoFoward1 struct {
}

func (p DepartGoFoward1) Name() string {
	return "DepartFoward1"
}

func (p DepartGoFoward1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if c.MapId() == 103040410 && script.QuestCompleted(l, c)(2287) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(103040420, "right01")
		return true
	}
	if c.MapId() == 103040420 && script.QuestCompleted(l, c)(2288) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(103040430, "right01")
		return true
	}
	if c.MapId() == 103040410 && script.QuestStarted(l, c)(2287) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(103040420, "right01")
		return true
	}
	if c.MapId() == 103040420 && script.QuestStarted(l, c)(2288) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(103040430, "right01")
		return true
	}
	if c.MapId() == 103040440 || c.MapId() == 103040450 {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(c.MapId()+10, "right01")
		return true
	}
	script.SendPinkNotice(l, c)("CANNOT_ACCESS")
	return false
}
