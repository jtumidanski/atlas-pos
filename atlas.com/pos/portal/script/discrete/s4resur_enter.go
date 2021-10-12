package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type S4ResurEnter struct {
}

func (p S4ResurEnter) Name() string {
	return "s4resur_enter"
}

func (p S4ResurEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.QuestStarted(l, c)(6134) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(922020000, 0)
		return true
	}
	script.SendPinkNotice(l, c)("MYSTERIOUS_FORCE")
	return false
}
