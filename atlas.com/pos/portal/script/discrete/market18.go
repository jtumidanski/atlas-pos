package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type Market18 struct {
}

func (p Market18) Name() string {
	return "market18"
}

func (p Market18) Enter(l logrus.FieldLogger, c script.Context) bool {
	return generic.EnterMarket(l, c)
}
