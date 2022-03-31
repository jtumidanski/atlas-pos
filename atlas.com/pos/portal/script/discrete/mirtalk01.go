package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MirTalk01 struct {
}

func (p MirTalk01) Name() string {
	return "mirtalk01"
}

func (p MirTalk01) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.BlockPortal(l, span, c)
	if processor.ContainsAreaInfo(l, c)(22013, "dt01=o") {
		return false
	}
	processor.MapEffect(l, c)("evan/dragonTalk01")
	processor.UpdateAreaInfo(l, c)(22013, "dt00=o;dt01=o;mo00=o;mo01=o;mo10=o;mo02=o")
	return true
}
