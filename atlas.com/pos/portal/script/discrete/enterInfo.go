package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterInfo struct {
}

func (p EnterInfo) Name() string {
	return "enterInfo"
}

func (p EnterInfo) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestActive(l, c)(21733) &&
		script.QuestProgressInt(l, c)(21733, 9300345) == 0 &&
		_map.MonstersCount(l)(c.WorldId(), c.ChannelId(), 104000004) == 0 {
		_map.SpawnMonster(l)(c.WorldId(), c.ChannelId(), 104000004, 9300345, 0, 0)
		script.SetQuestProgress(l, c)(21733, 21762, 2)
	}
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(104000004, 1)
	return true
}
