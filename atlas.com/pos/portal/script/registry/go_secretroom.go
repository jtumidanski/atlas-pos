package registry

import (
	"atlas-pos/portal/script"
	"github.com/sirupsen/logrus"
)

type GoSecretroom struct {
}

func (p GoSecretroom) Name() string {
	return "go_secretroom"
}

func (p GoSecretroom) Enter(l logrus.FieldLogger, c script.Context) bool {

}
