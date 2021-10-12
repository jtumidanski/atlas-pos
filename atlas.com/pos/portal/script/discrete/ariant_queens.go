package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantQueens struct {
}

func (p AriantQueens) Name() string {
	return "ariant_queens"
}

func (p AriantQueens) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.Morphed(l, c)(221005) {
		return false
	} else {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(260000300, 7)
		script.SendPinkNotice(l, c)("PALACE_INTRUDER")
		return true
	}
}
