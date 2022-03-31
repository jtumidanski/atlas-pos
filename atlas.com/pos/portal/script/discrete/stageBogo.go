package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type StageBogo struct {
}

func (p StageBogo) Name() string {
	return "stageB"
}

func (p StageBogo) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(670010800, 0)
	return true
}
