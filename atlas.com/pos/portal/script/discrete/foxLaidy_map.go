package discrete

import (
	"atlas-pos/character"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type FoxLaidyMap struct {
}

func (p FoxLaidyMap) Name() string {
	return "foxLaidy_map"
}

func (p FoxLaidyMap) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(3647) && processor.HasItem(l, c)(4031793) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(222010200, "east00")
		return true
	}

	if !processor.QuestStarted(l, c)(23647) {
		character.ForceStartQuest(l)(c.CharacterId(), 23647)
	}
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(922220000, "east00")
	return true
}
