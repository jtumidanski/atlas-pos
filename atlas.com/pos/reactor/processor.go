package reactor

import (
	"atlas-pos/model"
	"atlas-pos/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByNameModelProvider(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, reactorName string) model.Provider[Model] {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestByName(worldId, channelId, mapId, reactorName), makeModel)
	}
}

func ByName(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, reactorName string) (Model, error) {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) (Model, error) {
		return ByNameModelProvider(l, span)(worldId, channelId, mapId, reactorName)()
	}
}

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(id uint32) model.Provider[Model] {
	return func(id uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestById(id), makeModel)
	}
}

func GetById(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (Model, error) {
	return func(id uint32) (Model, error) {
		return ByIdModelProvider(l, span)(id)()
	}
}

func ForceHit(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorId uint32, state byte) {
	return func(worldId byte, channelId byte, mapId uint32, reactorId uint32, state byte) {

	}
}

func makeModel(data requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(data.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	attr := data.Attributes
	return Model{
		id:    uint32(id),
		state: attr.State,
	}, nil
}
