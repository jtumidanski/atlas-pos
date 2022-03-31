package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Skyrom struct {
}

func (p Skyrom) Name() string {
	return "skyrom"
}

func (p Skyrom) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestStarted(l, c)(3935) || processor.HasItem(l, c)(4031574) {
		return false
	}
	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 926000010) != 0 {
		processor.SendPinkNotice(l, c)("OTHER_PLAYER_TRYING")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(926000010, 0)
	return true
}
