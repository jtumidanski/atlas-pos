package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AranTutorOut1 struct {
}

func (p AranTutorOut1) Name() string {
	return "aranTutorOut1"
}

func (p AranTutorOut1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {

	if processor.QuestStarted(l, c)(21000) {
		processor.TeachSkill(l, c)(20000017, 0, -1, -1)
		processor.TeachSkill(l, c)(20000018, 0, -1, -1)
		processor.TeachSkill(l, c)(20000018, 1, 0, -1)
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(914000200, 1)
		return true
	} else {
		processor.SendPinkNotice(l, c)("ARAN_TUTORIAL_EXIT")
		return false
	}
}
