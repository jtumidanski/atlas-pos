package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantQueens struct {
}

func (p AriantQueens) Name() string {
	return "ariant_queens"
}

func (p AriantQueens) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.Morphed(l, c)(221005) {
		return false
	} else {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(260000300, 7)
		processor.SendPinkNotice(l, c)("PALACE_INTRUDER")
		return true
	}
}
