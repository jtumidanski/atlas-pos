package discrete

import (
	_map "atlas-pos/map"
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type ExitPuppeteer struct {
}

func (p ExitPuppeteer) Name() string {
	return "exit_puppeteer"
}

func (p ExitPuppeteer) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	if _map.MonsterCount(l)(c.WorldId(), c.ChannelId(), c.MapId(), 9300285) > 0 {
		processor.SendPinkNotice(l, c)("MUST_DEFEAT_PUPPETEER")
		return false
	}

	//EventInstanceManager eim = pi.getEventInstance()
	//if (eim != null) {
	//	eim.stopEventTimer()
	//	eim.dispose()
	//}
	processor.PlayPortalSound(l, c)
	processor.WarpById(l, span, c)(105070300, 3)
	return true
}
