package discrete

import (
	"atlas-pos/character"
	"atlas-pos/job"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EvanFarmCT struct {
}

func (p EvanFarmCT) Name() string {
	return "evanFarmCT"
}

func (p EvanFarmCT) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(22010) || !character.IsJob(l, span)(c.CharacterId(), job.Evan) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(100030310, 0)
	} else {
		script.SendPinkNotice(l, c)("CANNOT_ENTER_LUSH_FOREST_WITHOUT")
	}
	return true
}
