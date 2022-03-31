package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RienTutor8 struct {
}

func (p RienTutor8) Name() string {
	return "rienTutor8"
}

func (p RienTutor8) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if character.IsJob(l, span)(c.CharacterId(), job.Legend) {
		if processor.QuestStarted(l, c)(21015) {
			processor.ShowInfoText(l, c)("You must exit to the right in order to find Murupas.")
			return false
		}
		if processor.QuestStarted(l, c)(21016) {
			processor.ShowInfoText(l, c)("You must exit to the right in order to find Murupas.")
			return false
		}
		if processor.QuestStarted(l, c)(21017) {
			processor.ShowInfoText(l, c)("You must exit to the right in order to find Murupas.")
			return false
		}
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(140010000, 2)
	return true
}
