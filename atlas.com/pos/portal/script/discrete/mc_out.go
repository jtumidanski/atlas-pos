package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type McOut struct {
}

func (p McOut) Name() string {
	return "mc_out"
}

func (p McOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.PlayPortalSound(l, c)
	processor.WarpToSavedLocationDefaultMap(l, span, c)("MONSTER_CARNIVAL", 102000000)
	return true
}
