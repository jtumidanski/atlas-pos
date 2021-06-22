package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type ExitPuppeteer struct {
}

func (p ExitPuppeteer) Name() string {
	return "exit_puppeteer"
}

func (p ExitPuppeteer) Enter(l logrus.FieldLogger, c script.Context) bool {
	if _map.MonsterCount(l)(c.WorldId(), c.ChannelId(), c.MapId(), 9300285) > 0 {
		script.SendPinkNotice(l, c)("MUST_DEFEAT_PUPPETEER")
		return false
	}

	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	eim.stopEventTimer()
	//	eim.dispose()
	//}
	script.PlayPortalSound(l, c)
	script.WarpById(l, c)(105070300, 3)
	return true
}
