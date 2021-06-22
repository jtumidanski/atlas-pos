package discrete

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type Market00 struct {
}

func (p Market00) Name() string {
	return "market00"
}

func (p Market00) Enter(l logrus.FieldLogger, c script.Context) bool {
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(script.GetMarketPortal(l, c))
	return true
}
