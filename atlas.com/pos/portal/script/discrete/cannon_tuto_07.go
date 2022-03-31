package discrete

import (
	"atlas-pos/portal/script"
	"atlas-pos/portal/script/processor"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type CannonTuto07 struct {
}

func (p CannonTuto07) Name() string {
	return "cannon_tuto_07"
}

func (p CannonTuto07) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	processor.SetDirectionStatus(l, c)(true)
	processor.LockUI(l, c)
	processor.SpawnNPC(l, c)(579711, 1096012, -51, -97, 0, true)
	processor.SetNPCValue(l, c)(579711, "summon")
	processor.UpdateInfo(l, c)("fly", "579711")
	processor.SendDirectionInfo(l, c)(3, 0)
	processor.SendDirectionInfo(l, c)(3, 2)
	processor.SendDirectionInfo(l, c)(4, 0)
	return true
}
