package repository

import (
	"github.com/cyberpunkattack/api/base"
	"time"
)

type SessionRepository struct {
	*base.Repository
}

type NewSessionIM struct {
	Invites     []string
	Password    string
	Name        string
	CreatorHash string
}

type ActiveSession struct {
	Creator   string    `json:"creator"`
	SessionID string    `json:"session_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (session *SessionRepository) GetUserByEmail(email string) bool {
	return true
}

func (session *SessionRepository) CreateNewSessionIM(args NewSessionIM) (*ActiveSession, error) {

	return &ActiveSession{}, nil
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		Repository: &base.Repository{
			TableName: "sessions",
		},
	}
}
