package models

import "time"

type ExpirationModel struct {
	*BaseModel
	Type string `json:"type"`
	ExpireAfter time.Time `json:"expire_after"`
}