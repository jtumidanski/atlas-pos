package party

import (
	"atlas-pos/model"
	"atlas-pos/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByMemberIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(memberId uint32) model.Provider[Model] {
	return func(memberId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestByMemberId(memberId), makeModel)
	}
}

func GetByMemberId(l logrus.FieldLogger, span opentracing.Span) func(memberId uint32) (Model, error) {
	return func(memberId uint32) (Model, error) {
		return ByMemberIdModelProvider(l, span)(memberId)()
	}
}

func makeModel(body requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	m := Model{id: uint32(id)}
	return m, nil
}
