package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Resurrection struct {
}

func (p S4Resurrection) Name() string {
	return "s4resurrection"
}

func (p S4Resurrection) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(4001108) {
		if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), 923000100) == 0 {
			_map.ResetMapObjects(l)(c.WorldId(), c.ChannelId(), 923000100)
			processor.PlayPortalSound(l, c)
			processor.WarpById(l, span, c)(923000100, 0)
			return true
		} else {
			processor.SendPinkNotice(l, c)("OTHER_PLAYER_INSIDE")
			return false
		}
	}
	processor.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
