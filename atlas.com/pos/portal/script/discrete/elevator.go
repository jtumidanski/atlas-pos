package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Elevator struct {
}

func (p Elevator) Name() string {
	return "elevator"
}

func (p Elevator) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//EventManager elevator = pi.getEventManager("Elevator")
	//if (elevator == null) {
	//	pi.sendPinkNotice("ELEVATOR_MAINTENANCE")
	//} else if (elevator.getProperty(pi.getMapId() == 222020100 ? ("goingUp") : ("goingDown")) == "false") {
	//	pi.playPortalSound(); pi.warp(pi.getMapId() == 222020100 ? 222020110 : 222020210, 0)
	//	return true
	//} else if (elevator.getProperty(pi.getMapId() == 222020100 ? ("goingUp") : ("goingDown")) == "true") {
	//	pi.sendPinkNotice("ELEVATOR_MOVING")
	//}
	return false
}
