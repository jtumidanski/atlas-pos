package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type JailIn struct {
}

func (p JailIn) Name() string {
	return "jail_in"
}

func (p JailIn) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(300000012,"portal")
	return true
}
