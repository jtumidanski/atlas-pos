package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor2 struct {
}

func (p RienTutor2) Name() string {
	return "rienTutor2"
}

func (p RienTutor2) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !script.QuestCompleted(l, c)(21011) {
		script.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(140090300, 1)
	return true
}
