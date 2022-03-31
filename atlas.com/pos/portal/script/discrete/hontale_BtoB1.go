package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HorntaleBtoB1 struct {
}

func (p HorntaleBtoB1) Name() string {
	return "hontale_BtoB1"
}

func (p HorntaleBtoB1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), c.MapId()) == 1 {
		processor.SendLightBlueNotice(l, c)("HORNTAIL_LAST_PLAYER")
		return false
	}
	if processor.HasItem(l, c)(4001087) {
		processor.SendLightBlueNotice(l, c)("HORNTAIL_CANNOT_PASS_WITH_1ST_CRYSTAL")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(240050101, 0)
	return true
}
