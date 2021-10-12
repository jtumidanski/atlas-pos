package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDisguise3 struct {
}

func (p EnterDisguise3) Name() string {
	return "enterDisguise3"
}

func (p EnterDisguise3) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(20301) ||
		script.QuestStarted(l, c)(20302) ||
		script.QuestStarted(l, c)(20303) ||
		script.QuestStarted(l, c)(20304) ||
		script.QuestStarted(l, c)(20305) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010620) > 0 {
			script.SendPinkNotice(l, c)("SOMEONE_ALREADY_SEARCHING")
			return false
		}

		if script.HasItem(l, c)(4032103) {
			script.SendPinkNotice(l, c)("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
			return false
		}

		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(108010620, "out00")
		return true
	}

	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(130010110, "out00")
	return true
}
