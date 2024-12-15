package repository

import "github.com/cyberpunkattack/api/base"

type SessionRepository struct {
	*base.Repository
}

func (session *SessionRepository) GetUserByEmail(email string) bool {
	return true
}

func NewSessionRepository() *SessionRepository {
	return &SessionRepository{
		Repository: &base.Repository{
			TableName: "sessions",
		},
	}
}
