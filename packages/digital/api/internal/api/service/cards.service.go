package service

import (
	"types"

	_ "github.com/cyberpunkattack"
	"github.com/cyberpunkattack/api/base"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/helpers"
	"github.com/cyberpunkattack/prealloc"
)

type CardsService struct {
	*base.Service
}

func (cards *CardsService) GenerateRandomActionDeck() models.ActionDeckType {
	result := models.ActionDeckType{}
	cardsFile := prealloc.ActionsMap
	actionCounter := prealloc.CardsCounter

	for spec, count := range actionCounter.GetByDeck("actions") {
		duplicates := helpers.Duplicate(cardsFile[spec], count)

		result = append(result, duplicates...)
	}

	shuffledDeck := helpers.Shuffle(result)

	return shuffledDeck
}

func (cards *CardsService) GenerateRoles(playerCount int) []*types.RoleCardCommonType {
	counts := prealloc.RolesCounter.GetByCount(playerCount)

	rolesCards := prealloc.RolesMap
	result := []*types.RoleCardCommonType{}

	for spec, count := range counts {
		duplicated := helpers.Duplicate(rolesCards[spec], count)

		result = append(result, duplicated...)
	}

	return result
}

func (cards *CardsService) GenerateRandomImplantDeck() []*types.ImplantCardCommonType {
	result := models.ImplantDeckType{}

	count := prealloc.CardsCounter.Implants
	allImplants := prealloc.ImplantsMap

	for spec, count := range count {
		duplicate := helpers.Duplicate(allImplants[spec], count)

		result = append(result, duplicate...)
	}

	return result
}

func NewCardsService() *CardsService {
	return &CardsService{
		&base.Service{},
	}
}
