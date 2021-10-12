package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type JailOut struct {
}

func (p JailOut) Name() string {
	return "jail_out"
}

func (p JailOut) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//long jailedTime = pi.getJailTimeLeft()
	//
	//if (jailedTime <= 0) {
	//	pi.playPortalSound()
	//	pi.warp(300000010, "in01")
	//	return true
	//} else {
	//	int seconds = Math.floor(jailedTime / 1000) % 60
	//	int minutes = (Math.floor(jailedTime / (1000 * 60)) % 60)
	//	int hours = (Math.floor(jailedTime / (1000 * 60 * 60)) % 24)
	//
	//	pi.sendPinkNotice("JAIL_NOTE", hours, minutes, seconds)
		return false
	//}
}
