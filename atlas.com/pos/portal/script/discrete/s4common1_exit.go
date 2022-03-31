package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Common1Exit struct {
}

func (p S4Common1Exit) Name() string {
	return "s4common1_exit"
}

func (p S4Common1Exit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	if !processor.HasItem(l, c)(4031495) {
		processor.WarpRandom(l, span, c)(211040100)
		return true
	}
	processor.WarpRandom(l, span, c)(921100301)
	return true
}
