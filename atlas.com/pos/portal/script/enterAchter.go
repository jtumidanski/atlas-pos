package script

import "github.com/sirupsen/logrus"

type enterAchter struct {
}

func EnterAchter() PortalScript {
	return enterAchter{}
}

func (a enterAchter) Name() string {
	return "enterAchter"
}

func (a enterAchter) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	p.PlayPortalSound()
	p.WarpByName(100000201, "out02")
	return true
}
