package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DavyNext1 struct {
}

func (p DavyNext1) Name() string {
	return "davy_next1"
}

func (p DavyNext1) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if processor.GetEventProperty(l, c)("stage2") == "3" {
		processor.PlayPortalSound(l, c)
		processor.WarpById(l, span, c)(925100200, 0)
		return true
	}
	processor.SendPinkNotice(l, c)("PORTAL_NOT_YET_OPENED")
	return false
}
