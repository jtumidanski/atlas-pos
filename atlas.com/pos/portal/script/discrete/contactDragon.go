package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ContactDragon struct {
}

func (p ContactDragon) Name() string {
	return "contactDragon"
}

func (p ContactDragon) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(900090100, 0)
	return true
}
