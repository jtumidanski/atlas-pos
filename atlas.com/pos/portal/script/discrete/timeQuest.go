package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
)

type TimeQuest struct {
}

func (p TimeQuest) Name() string {
	return "timeQuest"
}

func (p TimeQuest) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	m := uint32(math.Floor(float64((c.MapId() - 270010000) / 100)))
	if m < 5 && processor.QuestCompleted(l, c)(3500+m) {
		processor.WarpByName(l, span, c)(c.MapId()+10, "out00")
		return true
	}
	if m == 5 && processor.QuestCompleted(l, c)(3502+m) {
		processor.WarpByName(l, span, c)(270020000, "out00")
		return true
	}
	if m > 100 && m < 105 && processor.QuestCompleted(l, c)(3407+m) {
		processor.WarpByName(l, span, c)(c.MapId()+10, "out00")
		return true
	}
	if m == 105 && processor.QuestCompleted(l, c)(3514) {
		processor.WarpByName(l, span, c)(270030000, "out00")
		return true
	}
	if m > 200 && m < 205 && processor.QuestCompleted(l, c)(3314+m) {
		processor.WarpByName(l, span, c)(c.MapId()+10, "out00")
		return true
	}
	if m == 205 && processor.QuestCompleted(l, c)(3519) {
		processor.WarpByName(l, span, c)(270040000, "out00")
		return true
	}
	if m == 300 && (processor.HasItem(l, c)(4032002) || processor.QuestCompleted(l, c)(3522)) {
		processor.WarpByName(l, span, c)(270040100, "out00")
		return true
	}
	processor.SendPinkNotice(l, c)("TIME_QUEST")
	if m > 200 {
		processor.WarpByName(l, span, c)(270030000, "in00")
		return true
	}
	if m > 100 {
		processor.WarpByName(l, span, c)(270020000, "in00")
		return true
	}
	processor.WarpByName(l, span, c)(270010000, "in00")
	return true
}
