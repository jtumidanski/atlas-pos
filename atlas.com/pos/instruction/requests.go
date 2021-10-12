package instruction

import (
	"atlas-pos/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	worldChannelPrefix      string = "/ms/csrv/"
	worldChannelService            = requests.BaseRequest + worldChannelPrefix
	worldsResource                 = worldChannelService + "worlds"
	worldResource                  = worldsResource + "/%d"
	channelsResource               = worldResource + "/channels"
	channelResource                = channelsResource + "/%d"
	worldCharactersResource        = channelResource + "/characters"
	worldCharacterResource         = worldCharactersResource + "/%d"
	instructionsResource           = worldCharacterResource + "/instructions"
)

func Create(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) error {
	return func(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) error {
		ar := &InputDataContainer{
			Data: DataBody{
				Id:   "0",
				Type: "Instruction",
				Attributes: Attributes{
					Message: message,
					Width:   width,
					Height:  height,
				},
			},
		}
		_, err := requests.Post(l, span)(fmt.Sprintf(instructionsResource, worldId, channelId, characterId), ar)
		if err != nil {
			return err
		}
		return nil
	}
}
