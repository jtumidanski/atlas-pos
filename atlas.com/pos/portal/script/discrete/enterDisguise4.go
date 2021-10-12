package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterDisguise4 struct {
}

func (p EnterDisguise4) Name() string {
	return "enterDisguise4"
}

func (p EnterDisguise4) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(20301) ||
		script.QuestStarted(l, c)(20302) ||
		script.QuestStarted(l, c)(20303) ||
		script.QuestStarted(l, c)(20304) ||
		script.QuestStarted(l, c)(20305) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 108010630) > 0 {
			script.SendPinkNotice(l, c)("SOMEONE_ALREADY_SEARCHING")
			return false
		}

		if script.HasItem(l, c)(4032104) {
			script.SendPinkNotice(l, c)("ALREADY_CHALLENGED_MASTER_OF_DISGUISE")
			return false
		}

		script.PlayPortalSound(l, c)
		script.WarpByName(l, span, c)(108010630, "out00")
		return true
	}

	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(130010120, "out00")
	return true
}
