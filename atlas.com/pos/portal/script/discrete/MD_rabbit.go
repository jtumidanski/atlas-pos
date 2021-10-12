package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDRabbit struct {
}

func (p MDRabbit) Name() string {
	return "MD_rabbit"
}

func (p MDRabbit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(221023400)
	dungeonId := uint32(221023401)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
