package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor4 struct {
}

func (p RienTutor4) Name() string {
	return "rienTutor4"
}

func (p RienTutor4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestCompleted(l, c)(21013) {
		script.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(140090500, 1)
	return true
}
