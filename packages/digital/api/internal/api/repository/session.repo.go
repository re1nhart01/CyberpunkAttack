package repository

import (
	"context"
	"strings"
	"time"

	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/crypto"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/mongo"
	"github.com/cyberpunkattack/environment"
	"github.com/cyberpunkattack/helpers"
)

type SessionRepository struct {
	*base.Repository
}

type NewSessionIM struct {
	Invites     []string
	Password    string
	Name        string
	CreatorHash string
	Type        string
}

type ActiveSession struct {
	Creator      string    `json:"creator"`
	SessionID    string    `json:"session_id"`
	CreatedAt    time.Time `json:"created_at"`
	PendingUsers []string  `json:"pending_users"`
	Name         string    `json:"name"`
}

func (session *SessionRepository) GetSessionByField(ctx context.Context, k string, v any) (*models.SessionIM, error) {
	result := &models.SessionIM{}
	sessionColl := mongo.DB().Get().Collection("sessions")

	if err := sessionColl.FindOne(ctx, map[string]any{k: v}).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

func (session *SessionRepository) AddUserToSession(ctx context.Context) (error) {


	return nil
}

func (session *SessionRepository) CreateNewCustomSessionIM(ctx context.Context, args NewSessionIM) (*ActiveSession, error) {
	sessionColl := mongo.DB().Get().Collection("sessions")
	hashedPassword := helpers.Ternary(args.Password == "", "", crypto.HashPassword(args.Password))
	sessionId := crypto.GetSha1(
		environment.GEnv().Get("SERVER_KEY"),
		strings.Join([]string{
			time.Now().String(),
			hashedPassword,
			args.Type,
		}, ":AT:")) + time.Now().String()

	createdSession := &models.SessionIM{
		Name:        args.Name,
		CreatorHash: args.CreatorHash,
		SessionID:   sessionId,
		Password:    hashedPassword,
		UserIds:     args.Invites,
		UserMap:     map[string]models.UserSessionModel{},
		ActionDeck:  []models.ActionDeckType{},
		MovesList:   []models.SessionMoveModel{},
		ImplantDeck: []models.ImplantDeckType{},
		Type:        models.GAME_TYPE_CUSTOM,
		IsEnded:     false,
		CreatedAt:   time.Now(),
		EndedAt:     time.Time{},
		Flags: models.SessionFlags{
			Started: false,
		},
	}

	if _, err := sessionColl.InsertOne(ctx, &createdSession); err != nil {
		return nil, err
	}

	activeSession := &ActiveSession{
		Creator:      args.CreatorHash,
		SessionID:    sessionId,
		CreatedAt:    time.Now(),
		Name:         args.Name,
		PendingUsers: args.Invites,
	}

	return activeSession, nil
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		Repository: &base.Repository{
			TableName: "sessions",
		},
	}
}
