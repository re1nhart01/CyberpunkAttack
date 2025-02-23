package gameplay

import (
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/prealloc"
	"time"
	"types"
)

// selector - file where you will select -> actions | cyberimplants -> which category -> run specific handler for every card category

type ActionObject struct {
	Creator   string         `json:"creator"`
	SessionId string         `json:"sessionId"`
	Deck      DECK           `json:"deck"`
	Card      string         `json:"card"`
	To        string         `json:"to"`
	From      string         `json:"from"`    //from your hand or from deck (from deck is for implants)
	Payload   map[string]any `json:"payload"` // if card needs something additional
}

type ExecutionObject struct {
	Creator   string           `json:"creator"`
	SessionId string           `json:"sessionId"`
	Session   models.SessionIM `json:"session"`
	Now       time.Time        `json:"now"`
	Action    ActionObject     `json:"action"`
}

func HandleSelectDeck(args *HandlersArgs) error {
	switch args.executionObj.Action.Deck {
	case ActionDeck:
		return HandleSelectActionCategory(args)
	case Implants:
		return HandleSelectImplantsCategory(args)
	default:
		return AbortGameAction(args)
	}
}

func HandleSelectActionCategory(args *HandlersArgs) error {
	selectedCard, ok := prealloc.ActionsMap[args.executionObj.Action.Card]
	if !ok {
		return AbortGameAction(args)
	}
	switch selectedCard.ActionType {
	case types.ActionTypeHacks:
		return HandleHacksCards(args)
	case types.ActionTypeBoom:
		return HandleBoomCards(args)
	case types.ActionTypeEvades:
		return HandleEvadesCards(args)
	case types.ActionTypeGun:
		return HandleGunCards(args)
	case types.ActionTypeHeals:
		return HandleHealsCards(args)
	case types.ActionTypeRunners:
		return HandleRunnersCards(args)
	case types.ActionTypeOther:
		return HandleOtherCards(args)
	default:
		return AbortGameAction(args)
	}
}

func HandleSelectImplantsCategory(args *HandlersArgs) error {
	selectedCard, ok := prealloc.ImplantsMap[args.executionObj.Action.Card]
	if !ok {
		return AbortGameAction(args)
	}
	switch selectedCard.ImplantType {
	case types.ImplantTypeBlood:
		return HandleBloodImplant(args)
	case types.ImplantTypeCybervirus:
		return HandleCybervirusImplant(args)
	case types.ImplantTypeBody:
		return HandleBodyImplant(args)
	case types.ImplantTypeNeural:
		return HandleNeuralImplant(args)
	case types.ImplantTypeFight:
		return HandleFightImplant(args)
	case types.ImplantTypeSkeleton:
		return HandleSkeletonImplant(args)
	case types.ImplantTypeOther:
		return HandleOtherImplant(args)
	default:
		return AbortGameAction(args)
	}
}
