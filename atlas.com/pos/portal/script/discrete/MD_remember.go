package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDRemember struct {
}

func (p MDRemember) Name() string {
	return "MD_remember"
}

func (p MDRemember) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(240040511)
	dungeonId := uint32(240040800)
	dungeons := uint8(19)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
