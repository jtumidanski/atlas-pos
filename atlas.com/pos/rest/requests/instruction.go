package requests

import (
	"atlas-pos/rest/attributes"
	"fmt"
)

const (
	worldChannelPrefix      string = "/ms/csrv/"
	worldChannelService            = BaseRequest + worldChannelPrefix
	worldsResource                 = worldChannelService + "worlds"
	worldResource                  = worldsResource + "/%d"
	channelsResource               = worldResource + "/channels"
	channelResource                = channelsResource + "/%d"
	worldCharactersResource        = channelResource + "/characters"
	worldCharacterResource         = worldCharactersResource + "/%d"
	instructionsResource           = worldCharacterResource + "/instructions"
)

var WorldChannel = func() *worldChannel {
	return &worldChannel{}
}

type worldChannel struct {
}

func (c *worldChannel) CreateInstruction(worldId byte, channelId byte, characterId uint32, message string, width int16, height int16) error {
	ar := &attributes.InstructionInputDataContainer{
		Data: attributes.InstructionData{
			Id:   "0",
			Type: "Instruction",
			Attributes: attributes.InstructionAttributes{
				Message: message,
				Width:   width,
				Height:  height,
			},
		},
	}
	_, err := Post(fmt.Sprintf(instructionsResource, worldId, channelId, characterId), ar)
	if err != nil {
		return err
	}
	return nil
}
