package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type DepartGoFoward0 struct {
}

func (p DepartGoFoward0) Name() string {
	return "DepartFoward0"
}

func (p DepartGoFoward0) Enter(l logrus.FieldLogger, c script.Context) bool {
	if c.MapId() == 103040410 && script.QuestCompleted(l, c)(2287) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(103040420, "right00")
		return true
	}
	if c.MapId() == 103040420 && script.QuestCompleted(l, c)(2288) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(103040430, "right00")
		return true
	}
	if c.MapId() == 103040410 && script.QuestStarted(l, c)(2287) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(103040420, "right00")
		return true
	}
	if c.MapId() == 103040420 && script.QuestStarted(l, c)(2288) {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(103040430, "right00")
		return true
	}
	if c.MapId() == 103040440 || c.MapId() == 103040450 {
		script.PlayPortalSound(l, c)
		script.WarpByName(l, c)(c.MapId()+10, "right00")
		return true
	}
	script.SendPinkNotice(l, c)("CANNOT_ACCESS")
	return false
}
