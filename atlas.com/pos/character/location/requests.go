package location

import (
	"atlas-pos/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	charactersServicePrefix   string = "/ms/cos/"
	charactersService                = requests.BaseRequest + charactersServicePrefix
	charactersResource               = charactersService + "characters/"
	charactersById                   = charactersResource + "%d"
	charactersLocations              = charactersById + "/locations"
	charactersLocationsByType        = charactersLocations + "?type=%s"
)

func getLocation(characterId uint32, location string) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(charactersLocationsByType, characterId, location))
}

func saveLocation(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, location string) (requests.DataBody[attributes], error) {
	return func(characterId uint32, location string) (requests.DataBody[attributes], error) {
		i := inputDataContainer{
			Data: inputDataBody{
				Id:   "0",
				Type: "com.atlas.cos.rest.attribute.CharacterSeedAttributes",
				Attributes: inputAttributes{
					Type: location,
				},
			},
		}
		r, _, err := requests.MakePostRequest[attributes](fmt.Sprintf(charactersLocations, characterId), i)(l, span)
		return r.Data(), err
	}
}
