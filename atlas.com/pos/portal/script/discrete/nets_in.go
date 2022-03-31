package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type NetsIn struct {
}

func (p NetsIn) Name() string {
	return "nets_in"
}

func (p NetsIn) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.SaveLocation(l, c)("MIRROR")
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(926010000, 4)
	return true
}
