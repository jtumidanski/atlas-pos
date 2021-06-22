package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/generic"
	"github.com/sirupsen/logrus"
)

type Market54 struct {
}

func (p Market54) Name() string {
	return "market54"
}

func (p Market54) Enter(l logrus.FieldLogger, c script.Context) bool {
	return generic.EnterMarket(l, c)
}
