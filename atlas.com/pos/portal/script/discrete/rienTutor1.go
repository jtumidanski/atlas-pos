package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor1 struct {
}

func (p RienTutor1) Name() string {
	return "rienTutor1"
}

func (p RienTutor1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestCompleted(l, c)(21010) {
		processor.SendPinkNotice(l, c)("COMPLETE_QUEST_BEFORE_PROCEEDING")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(140090200, 1)
	return true
}
