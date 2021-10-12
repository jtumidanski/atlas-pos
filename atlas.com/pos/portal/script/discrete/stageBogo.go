package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type StageBogo struct {
}

func (p StageBogo) Name() string {
	return "stageB"
}

func (p StageBogo) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(670010800, 0)
	return true
}
