package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type Market26 struct {
}

func (p Market26) Name() string {
	return "market26"
}

func (p Market26) Enter(l logrus.FieldLogger, c script.Context) bool {
	return generic.EnterMarket(l, c)
}
