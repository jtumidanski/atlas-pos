package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AriantCastle struct {
}

func (p AriantCastle) Name() string {
	return "ariant_castle"
}

func (p AriantCastle) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.HasItem(l, c)(4031582) {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(260000301, 5)
		return true
	} else {
		processor.SendPinkNotice(l, c)("ENTRY_PASS_NEEDED")
		return false
	}
}
