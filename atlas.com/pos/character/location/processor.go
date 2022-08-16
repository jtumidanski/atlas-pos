package location

import (
	"atlas-pos/model"
	"atlas-pos/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func SaveLocation(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, location string) {
	return func(characterId uint32, location string) {
		_, err := saveLocation(l, span)(characterId, location)
		if err != nil {
			l.WithError(err).Errorf("Unable to save location %s for character %d.", location, characterId)
			return
		}
	}
}

func ByTypeProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, location string) model.SliceProvider[Model] {
	return func(characterId uint32, location string) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(getLocation(characterId, location), makeLocation)
	}
}

func GetLocation(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, location string) (Model, error) {
	return func(characterId uint32, location string) (Model, error) {
		return model.SliceProviderToProviderAdapter(ByTypeProvider(l, span)(characterId, location), model.RandomPreciselyOneFilter[Model])()
	}
}

func makeLocation(b requests.DataBody[attributes]) (Model, error) {
	return Model{
		locationType: b.Attributes.Type,
		mapId:        b.Attributes.MapId,
		portalId:     b.Attributes.PortalId,
	}, nil
}
