package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MDError struct {
}

func (p MDError) Name() string {
	return "MD_error"
}

func (p MDError) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	baseId := uint32(261020300)
	dungeonId := uint32(261020301)
	dungeons := uint8(30)
	return generic.EnterMiniDungeon(l, span, c)(baseId, dungeonId, dungeons)
}
