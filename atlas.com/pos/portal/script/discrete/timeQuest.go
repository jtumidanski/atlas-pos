package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
	"math"
)

type TimeQuest struct {
}

func (p TimeQuest) Name() string {
	return "timeQuest"
}

func (p TimeQuest) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	m := uint32(math.Floor(float64((c.MapId() - 270010000) / 100)))
	if m < 5 && script.QuestCompleted(l, c)(3500+m) {
		script.WarpByName(l, c)(c.MapId()+10, "out00")
		return true
	}
	if m == 5 && script.QuestCompleted(l, c)(3502+m) {
		script.WarpByName(l, c)(270020000, "out00")
		return true
	}
	if m > 100 && m < 105 && script.QuestCompleted(l, c)(3407+m) {
		script.WarpByName(l, c)(c.MapId()+10, "out00")
		return true
	}
	if m == 105 && script.QuestCompleted(l, c)(3514) {
		script.WarpByName(l, c)(270030000, "out00")
		return true
	}
	if m > 200 && m < 205 && script.QuestCompleted(l, c)(3314+m) {
		script.WarpByName(l, c)(c.MapId()+10, "out00")
		return true
	}
	if m == 205 && script.QuestCompleted(l, c)(3519) {
		script.WarpByName(l, c)(270040000, "out00")
		return true
	}
	if m == 300 && (script.HasItem(l, c)(4032002) || script.QuestCompleted(l, c)(3522)) {
		script.WarpByName(l, c)(270040100, "out00")
		return true
	}
	script.SendPinkNotice(l, c)("TIME_QUEST")
	if m > 200 {
		script.WarpByName(l, c)(270030000, "in00")
		return true
	}
	if m > 100 {
		script.WarpByName(l, c)(270020000, "in00")
		return true
	}
	script.WarpByName(l, c)(270010000, "in00")
	return true
}
