package script

import (
	"github.com/sirupsen/logrus"
)

type apq2 struct {
}

func Apq2() PortalScript {
	return apq2{}
}

func (a apq2) Name() string {
	return "apq2"
}

func (a apq2) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpById(670010500, 0)
	return true
}