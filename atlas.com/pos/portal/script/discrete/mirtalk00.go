package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MirTalk00 struct {
}

func (p MirTalk00) Name() string {
	return "mirtalk00"
}

func (p MirTalk00) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "dt00=o") {
		return false
	}
	processor.MapEffect(l, c)("evan/dragonTalk00")
	processor.UpdateAreaInfo(l, c)(22013, "dt00=o;mo00=o")
	return true
}
