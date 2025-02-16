package service

import (
	"context"
	"errors"
	"github.com/cyberpunkattack/api/base"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/mongo"
	"github.com/cyberpunkattack/pkg/dispatcher"
)

type dict map[string]any

type GameService struct {
	*base.Service
}

func (game *GameService) CommitSessionUpdate(
	ctx *context.Context,
	sessionId string,
	tx func(payload models.SessionCommitment) (models.SessionCommitment, error),
) error {
	return nil
}

func (game *GameService) GetAllSubscriptions(ctx context.Context) dispatcher.DispatcherSubscription {
	// ctx := context.Background()
	list := dispatcher.DispatcherSubscription{}
	return list
}

func (game *GameService) StartGame(ctx context.Context, sessionId string, creatorHash string) error {
	sessionColl := mongo.DB().Get().Collection("sessions")

	session := models.SessionIM{}
	if err := sessionColl.FindOne(ctx, dict{"sessionId": sessionId}).Decode(&session); err != nil {
		return errors.New("session with this id is not exists")
	}

	if session.Status != models.PENDING_STATUS {
		return errors.New("session could be already started")
	}

	if session.CreatorHash == creatorHash {
		return errors.New("session's creator is not same")
	}

	actionDeck := models.ActionDeckType{}
	implantsDeck := models.ImplantDeckType{}

	return nil
}

func (game *GameService) UnsubscribeAllEvents(ctx context.Context) error {
	return nil
}

func NewGameService() *GameService {
	return &GameService{
		&base.Service{},
	}
}
