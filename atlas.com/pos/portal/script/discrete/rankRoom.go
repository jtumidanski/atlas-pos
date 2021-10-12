package discrete

import (
	"atlas-pos/portal/script"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type RankRoom struct {
}

func (p RankRoom) Name() string {
	return "rankRoom"
}

func (p RankRoom) Enter(l logrus.FieldLogger, span opentracing.Span, c script.Context) bool {
	script.PlayPortalSound(l, c)

	switch c.MapId() {
	case 130000000:
		script.WarpById(l, span, c)(130000100, 5) //or 130000101
		break
	case 130000200:
		script.WarpById(l, span, c)(130000100, 4) //or 130000101
		break
	case 140010100:
		script.WarpById(l, span, c)(140010110, 1) //or 140010111
		break
	case 120000101:
		script.WarpById(l, span, c)(120000105, 1)
		break
	case 103000003:
		script.WarpById(l, span, c)(103000008, 1) //or 103000009
		break
	case 100000201:
		script.WarpById(l, span, c)(100000204, 2) //or 100000205
		break
	case 101000003:
		script.WarpById(l, span, c)(101000004, 2) //or 101000005
		break
	default:
		script.WarpById(l, span, c)(c.MapId()+1, 1) //or + 2
		break
	}

	return true
}
