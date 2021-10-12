package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type LionCastleEnter struct {
}

func (p LionCastleEnter) Name() string {
	return "lionCastle_enter"
}

func (p LionCastleEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpByName(l, span, c)(211060010, "west00")
	return true
}
