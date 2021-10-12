package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type GuildWaitingEnter struct {
}

func (p GuildWaitingEnter) Name() string {
	return "guildwaitingenter"
}

func (p GuildWaitingEnter) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	//long entryTime = pi.getPlayer().getEventInstance().getProperty("entryTimestamp").toLong()
	//long timeNow = System.currentTimeMillis()
	//
	//int timeLeft = Math.ceil((entryTime - timeNow) / 1000).toInteger()
	//
	//if(timeLeft <= 0) {
	//	pi.playPortalSound(); pi.warp(990000100, 0)
	//	return true
	//}
	//else { //cannot proceed while allies can still enter
	//	pi.sendPinkNotice("PORTAL_OPEN_IN", timeLeft)
	//	return false
	//}
	return false
}
