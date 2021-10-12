package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDProtect struct {
}

func (p MDProtect) Name() string {
	return "MD_protect"
}

func (p MDProtect) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(240040520)
	dungeonId := uint32(240040900)
	dungeons := uint8(19)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
