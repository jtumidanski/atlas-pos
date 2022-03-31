package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GotoCastle struct {
}

func (p GotoCastle) Name() string {
	return "gotocastle"
}

func (p GotoCastle) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(2324) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(106020501, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("NEED_THORN_REMOVER")
	return false
}
