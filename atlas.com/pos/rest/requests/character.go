package requests

import (
	"atlas-pos/rest/attributes"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = baseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters"
	charactersById                 = charactersResource + "/%d"
	accountCharacters              = charactersService + "?accountId=%d&worldId=%d"
)

var Character = func() *character {
	return &character{}
}

type character struct {
}

func (c *character) GetCharacterAttributesById(characterId uint32) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := get(fmt.Sprintf(charactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func (c *character) GetAccountCharacters(accountId uint32, worldId byte) (*attributes.CharacterAttributesDataContainer, error) {
	ar := &attributes.CharacterAttributesDataContainer{}
	err := get(fmt.Sprintf(accountCharacters, accountId, worldId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
