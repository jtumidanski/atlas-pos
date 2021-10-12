package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDMushroom struct {
}

func (p MDMushroom) Name() string {
	return "MD_mushroom"
}

func (p MDMushroom) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(105050100)
	dungeonId := uint32(105050101)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
