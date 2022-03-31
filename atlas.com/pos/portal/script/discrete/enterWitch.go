package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterWitch struct {
}

func (p EnterWitch) Name() string {
	return "enterWitch"
}

func (p EnterWitch) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if !processor.QuestCompleted(l, c)(20404) {
		processor.SendPinkNotice(l, c)("SHOULD_NOT_GO_CREEPY")
		return false
	}

	mapId := uint32(924010000)
	if processor.QuestCompleted(l, c)(20407) {
		mapId = 924010200
	} else if processor.QuestCompleted(l, c)(20406) {
		mapId = 924010100
	}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(mapId, 1)
	return true
}
