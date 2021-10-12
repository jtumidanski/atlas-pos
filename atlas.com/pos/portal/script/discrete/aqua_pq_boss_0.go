package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type AquaPQBoss0 struct {
}

func (p AquaPQBoss0) Name() string {
	return "aqua_pq_boss_0"
}

func (p AquaPQBoss0) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(230040420, 0)
	return true
}
