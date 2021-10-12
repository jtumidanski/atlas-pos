package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type HorntaleBOpen struct {
}

func (p HorntaleBOpen) Name() string {
	return "hontale_Bopen"
}

func (p HorntaleBOpen) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//var nextMap uint32

	if c.MapId() == 240050101 {
		//nextMap = 240050102
		//eim = pi.getPlayer().getEventInstance()
		//target = eim.getMapInstance(nextMap)
		//targetPortal = target.getPortal("sp")
		//// only let people through if the eim is ready
		//avail = eim.getProperty("1stageclear")
		//if (avail == null) {
		//	// do nothing; send message to player
		script.SendLightBlueNotice(l, c)("HORNTAIL_SEAL")
		return false
		//} else {
		//	pi.playPortalSound()
		//	pi.getPlayer().changeMap(target, targetPortal)
		//	return true
		//}
	} else if c.MapId() == 240050102 {
		//nextMap = 240050103
		//eim = pi.getPlayer().getEventInstance()
		//target = eim.getMapInstance(nextMap)
		//targetPortal = target.getPortal("sp")
		//// only let people through if the eim is ready
		//avail = eim.getProperty("2stageclear")
		//if avail == null {
		//	// do nothing; send message to player
		script.SendLightBlueNotice(l, c)("HORNTAIL_SEAL")
		return false
		//} else {
		//	pi.playPortalSound()
		//	pi.getPlayer().changeMap(target, targetPortal)
		//	return true
		//}
	} else if c.MapId() == 240050103 {
		//nextMap = 240050104
		//eim = pi.getPlayer().getEventInstance()
		//target = eim.getMapInstance(nextMap)
		//targetPortal = target.getPortal("sp")
		//// only let people through if the eim is ready
		//avail = eim.getProperty("3stageclear")
		//if avail == null {
		// do nothing; send message to player
		script.SendLightBlueNotice(l, c)("HORNTAIL_SEAL")
		return false
		//} else {
		//	pi.playPortalSound()
		//	pi.getPlayer().changeMap(target, targetPortal)
		//	return true
		//}
	} else if c.MapId() == 240050104 {
		//nextMap = 240050105
		//eim = pi.getPlayer().getEventInstance()
		//target = eim.getMapInstance(nextMap)
		//targetPortal = target.getPortal("sp")
		//// only let people through if the eim is ready
		//avail = eim.getProperty("4stageclear")
		//if avail == null {
		// do nothing; send message to player
		script.SendLightBlueNotice(l, c)("HORNTAIL_SEAL")
		return false
		//} else {
		//	pi.playPortalSound()
		//	pi.getPlayer().changeMap(target, targetPortal)
		//	return true
		//}
	} else if c.MapId() == 240050105 {
		//nextMap = 240050100
		//eim = pi.getPlayer().getEventInstance()
		//target = eim.getMapInstance(nextMap)
		//targetPortal = target.getPortal("st00")
		//
		//avail = eim.getProperty("5stageclear")
		//if avail == null {
		//	if pi.haveItem(4001092) && pi.isEventLeader() {
		//		eim.showClearEffect()
		//		MessageBroadcaster.getInstance().sendServerNotice(pi.getPlayer(), ServerNoticeType.LIGHT_BLUE, I18nMessage.from("HORNTAIL_SEAL_BROKEN"))
		//		pi.playPortalSound()
		//		pi.getPlayer().changeMap(target, targetPortal)
		//		eim.setIntProperty("5stageclear", 1)
		//		return true
		//	} else {
		script.SendLightBlueNotice(l, c)("HORNTAIL_SEAL_LONG")
		return false
	}
	//} else {
	//	pi.playPortalSound()
	//	pi.getPlayer().changeMap(target, targetPortal)
	//	return true
	//}
	//}
	return true
}
