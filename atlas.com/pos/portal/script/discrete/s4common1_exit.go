package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4Common1Exit struct {
}

func (p S4Common1Exit) Name() string {
	return "s4common1_exit"
}

func (p S4Common1Exit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	if !script.HasItem(l, c)(4031495) {
		script.WarpRandom(l, span, c)(211040100)
		return true
	}
	script.WarpRandom(l, span, c)(921100301)
	return true
}
