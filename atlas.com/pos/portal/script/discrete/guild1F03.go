package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Guild1F03 struct {
}

func (p Guild1F03) Name() string {
	return "guild1F03"
}

func (p Guild1F03) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//pi.getEventInstance().gridInsert(pi.getPlayer(), 3)
	processor.PlayPortalSound(l, c)
	processor.WarpByName(l, span, c)(990000700, "st00")
	return true
}
