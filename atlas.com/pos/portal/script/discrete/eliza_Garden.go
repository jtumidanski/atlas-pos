package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ElizaGarden struct {
}

func (p ElizaGarden) Name() string {
	return "eliza_Garden"
}

func (p ElizaGarden) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(920020000, 2)
	return true
}
