package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type OutMagicLib struct {
}

func (p OutMagicLib) Name() string {
	return "outMagiclib"
}

func (p OutMagicLib) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if _map.MonsterCount(l)(c.WorldId(), c.ChannelId(), c.MapId(), 2220100) > 0 {
		processor.SendPinkNotice(l, c)("DEFEAT_BLUE_MUSHROOMS")
		return false
	}
	//EventInstanceManager eim = pi.getEventInstance()
	//eim.stopEventTimer()
	//eim.dispose()
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(101000000, 26)
	if processor.QuestCompleted(l, c)(20718) {
		processor.OpenNPCWithScript(l, c)(1103003, "MaybeItsGrendel_end")
	}
	return true
}
