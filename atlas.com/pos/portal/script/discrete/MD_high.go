package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDHigh struct {
}

func (p MDHigh) Name() string {
	return "MD_high"
}

func (p MDHigh) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(551030000)
	dungeonId := uint32(551030001)
	dungeons := uint8(19)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
