package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type Market01 struct {
}

func (p Market01) Name() string {
	return "market01"
}

func (p Market01) Enter(l logrus.FieldLogger, c script.Context) bool {
	return generic.EnterMarket(l, c)
}
