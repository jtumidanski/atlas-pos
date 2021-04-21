package script

import (
	"github.com/sirupsen/logrus"
)

type adviceMap struct {
}

func AdviceMap() PortalScript {
	return adviceMap{}
}

func (a adviceMap) Name() string {
	return "adviceMap"
}

func (a adviceMap) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.ShowInstruction("Press the #e#b[Up]#k arrow#n to use the portal and move to the next map.", 230, 5)
	return true
}
