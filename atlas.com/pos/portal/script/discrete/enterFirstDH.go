package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterFirstDH struct {
}

func (p EnterFirstDH) Name() string {
	return "enterFirstDH"
}

func (p EnterFirstDH) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	mapId := uint32(0)
	if processor.QuestStarted(l, c)(20701) {
		mapId = 913000000
	} else if processor.QuestStarted(l, c)(20702) {
		mapId = 913000100
	} else if processor.QuestStarted(l, c)(20703) {
		mapId = 913000200
	}

	if mapId == 0 {
		processor.SendPinkNotice(l, c)("KIKUS_ACCLIMATION_TRAINING_REQUIREMENT")
		return false
	}

	if _map.CharacterCount(l)(c.WorldId(), c.ChannelId(), mapId) > 0 {
		processor.SendPinkNotice(l, c)("SOMEONE_ALREADY_IN_MAP")
		return false
	}

	_map.ResetPartyQuest(l)(c.WorldId(), c.ChannelId(), mapId)
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(mapId, 0)
	return true
}
