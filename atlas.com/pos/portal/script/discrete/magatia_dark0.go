package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MagatiaDark0 struct {
}

func (p MagatiaDark0) Name() string {
	return "magatia_dark0"
}

func (p MagatiaDark0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.QuestCompleted(l, c)(7770) {
		processor.PlayPortalSound(l, c)
		processor.WarpByName(l, span, c)(926130000, "out00")
		return true
	}
	processor.SendPinkNotice(l, c)("PIPE_TOO_DARK")
	return false
}
