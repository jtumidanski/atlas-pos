package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDSand struct {
}

func (p MDSand) Name() string {
	return "MD_sand"
}

func (p MDSand) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(260020600)
	dungeonId := uint32(260020630)
	dungeons := uint8(34)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
