package service

import (
	_ "github.com/cyberpunkattack"
	"github.com/cyberpunkattack/api/base"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/type"
)

type CardsService struct {
	*base.Service
}

func (cards *CardsService) GenerateRandomActionDeck(deck *models.ActionDeckType) error {
	result := models.ActionDeckType{}
	//embed.EmbedActionsDeck
	cardsFile := Parsed

	return nil
}

func (cards *CardsService) GenerateRandomImplantDeck(deck *models.ImplantDeckType) error {

	return nil
}

func NewCardsService() *CardsService {
	return &CardsService{
		&base.Service{},
	}
}
