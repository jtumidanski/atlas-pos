package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDPig struct {
}

func (p MDPig) Name() string {
	return "MD_pig"
}

func (p MDPig) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(100020000)
	dungeonId := uint32(100020100)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
