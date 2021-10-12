package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantCastle struct {
}

func (p AriantCastle) Name() string {
	return "ariant_castle"
}

func (p AriantCastle) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if script.HasItem(l, c)(4031582) {
		script.PlayPortalSound(l, c)
		script.WarpById(l, span, c)(260000301, 5)
		return true
	} else {
		script.SendPinkNotice(l, c)("ENTRY_PASS_NEEDED")
		return false
	}
}
