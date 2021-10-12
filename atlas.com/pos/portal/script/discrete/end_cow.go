package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EndCow struct {
}

func (p EndCow) Name() string {
	return "end_cow"
}

func (p EndCow) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(2180) && (script.HasItem(l, c)(4031847) || script.HasItem(l, c)(4031848) || script.HasItem(l, c)(4031849) || script.HasItem(l, c)(4031850)) {
		if script.HasItem(l, c)(4031850) {
			script.PlayPortalSound(l, c)
			script.WarpRandom(l, span, c)(120000103)
			return true
		}
		script.SendPinkNotice(l, c)("COW_END")
		return false
	}
	script.PlayPortalSound(l, c)
	script.WarpRandom(l, span, c)(120000103)
	return true
}
