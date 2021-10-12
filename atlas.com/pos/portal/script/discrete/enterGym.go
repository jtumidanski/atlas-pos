package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterGym struct {
}

func (p EnterGym) Name() string {
	return "enterGym"
}

func (p EnterGym) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(21701) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(914010000, 1)
		return true
	}
	if script.QuestStarted(l, c)(21702) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(914010100, 1)
		return true
	}
	if script.QuestStarted(l, c)(21703) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(914010200, 1)
		return true
	}
	script.SendPinkNotice(l, c)("PUO_LESSON_REQUIREMENT")
	return false
}
