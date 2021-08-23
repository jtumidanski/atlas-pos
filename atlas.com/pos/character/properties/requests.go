package properties

import (
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

func request(l logrus.FieldLogger) func(url string) (*CharacterAttributesDataContainer, error) {
	return func(url string) (*CharacterAttributesDataContainer, error) {
		ar := &CharacterAttributesDataContainer{}
		err := requests.Get(l)(url, ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestAttributesById(l logrus.FieldLogger) func(characterId uint32) (*CharacterAttributesDataContainer, error) {
	return func(characterId uint32) (*CharacterAttributesDataContainer, error) {
		return request(l)(fmt.Sprintf(charactersById, characterId))
	}
}

func requestAccountCharacters(l logrus.FieldLogger) func(accountId uint32, worldId byte) (*CharacterAttributesDataContainer, error) {
	return func(accountId uint32, worldId byte) (*CharacterAttributesDataContainer, error) {
		return request(l)(fmt.Sprintf(accountCharacters, accountId, worldId))
	}
}
