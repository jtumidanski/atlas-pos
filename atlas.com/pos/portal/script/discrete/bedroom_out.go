package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type BedroomOut struct {
}

func (p BedroomOut) Name() string {
	return "bedroom_out"
}

func (p BedroomOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(2570) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(120000101, 0)
		return true
	}
	script.EarnTitle(l, c)("You still got some stuff to take care of. I can see it in your eyes. Wait...no, those are eye boogers.")
	return false
}
