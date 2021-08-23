package character

import (
	"atlas-pos/rest/attributes"
	"atlas-pos/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters"
	charactersById                 = charactersResource + "/%d"
	accountCharacters              = charactersService + "?accountId=%d&worldId=%d"
)

func requestAttributesById(l logrus.FieldLogger) func(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	return func(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
		ar := &attributes.CharacterAttributesDataContainer{}
		err := requests.Get(l)(fmt.Sprintf(charactersById, characterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestAccountCharacters(l logrus.FieldLogger) func(accountId uint32, worldId byte) (*attributes.CharacterAttributesDataContainer, error) {
	return func(accountId uint32, worldId byte) (*attributes.CharacterAttributesDataContainer, error) {
		ar := &attributes.CharacterAttributesDataContainer{}
		err := requests.Get(l)(fmt.Sprintf(accountCharacters, accountId, worldId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
