package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EndCow struct {
}

func (p EndCow) Name() string {
	return "end_cow"
}

func (p EndCow) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestStarted(l, c)(2180) && (processor.HasItem(l, c)(4031847) || processor.HasItem(l, c)(4031848) || processor.HasItem(l, c)(4031849) || processor.HasItem(l, c)(4031850)) {
		if processor.HasItem(l, c)(4031850) {
			processor.PlayPortalSound(l, c)
			processor.WarpRandom(l, span, c)(120000103)
			return true
		}
		processor.SendPinkNotice(l, c)("COW_END")
		return false
	}
	processor.PlayPortalSound(l, c)
	processor.WarpRandom(l, span, c)(120000103)
	return true
}
