package script

import "github.com/sirupsen/logrus"

type enterMagiclibrar struct {
}

func EnterMagiclibrar() PortalScript {
	return enterMagiclibrar{}
}

func (a enterMagiclibrar) Name() string {
	return "enterMagiclibrar"
}

func (a enterMagiclibrar) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.QuestStarted(20718) {

	} else {
		p.PlayPortalSound()
		p.WarpById(101000003, 8)
	}
	return true
}
