package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RnjExit struct {
}

func (p RnjExit) Name() string {
	return "rnj_exit"
}

func (p RnjExit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(261000011, 0)
	return true
}
