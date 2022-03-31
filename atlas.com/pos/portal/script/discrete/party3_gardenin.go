package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Party3GardenIn struct {
}

func (p Party3GardenIn) Name() string {
	return "party3_gardenin"
}

func (p Party3GardenIn) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//if (pi.getParty().isPresent() && pi.isEventLeader() && pi.hasItem(4001055, 1)) {
	//	pi.playPortalSound()
	//	pi.getEventInstance().warpEventTeam(920010100)
	//	return true
	//} else {
	processor.SendPinkNotice(l, c)("NEED_TO_BE_LEADER_AND_HAVE_ROOT_OF_LIFE")
	return false
	//}
}
