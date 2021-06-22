package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type Market21 struct {
}

func (p Market21) Name() string {
	return "market21"
}

func (p Market21) Enter(l logrus.FieldLogger, c script.Context) bool {
	return generic.EnterMarket(l, c)
}
