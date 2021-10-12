package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type EnterWarehouse struct {
}

func (p EnterWarehouse) Name() string {
	return "enterWarehouse"
}

func (p EnterWarehouse) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, span, c)(300000011, 0)
	return true
}
