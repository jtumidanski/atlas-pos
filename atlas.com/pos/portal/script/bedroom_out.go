package script

import (
	"github.com/sirupsen/logrus"
)

type bedroomOut struct {
}

func BedroomOut() PortalScript {
	return bedroomOut{}
}

func (a bedroomOut) Name() string {
	return "bedroom_out"
}

func (a bedroomOut) Enter(l logrus.FieldLogger, context Context) bool {
	p := Processor(l, context)
	if p.QuestStarted(2570) {
		p.PlayPortalSound()
		p.WarpById(120000101, 0)
		return true
	}
	p.EarnTitle("You still got some stuff to take care of. I can see it in your eyes. Wait...no, those are eye boogers.")
	return false
}
