package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type DojangExit struct {
}

func (p DojangExit) Name() string {
	return "dojang_exit"
}

func (p DojangExit) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpToSavedLocationDefaultMap(l, span, c)("MIRROR", 100000000)
	return true
}
