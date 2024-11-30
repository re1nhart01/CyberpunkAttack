package models

import (
    "fmt"
    "github.com/cyberpunkattack/database"
    "time"
)

type PreviousDataModel struct {
	BaseModel
	Type string `json:"type"`
	Value string `json:"value"`
	UserHashId string `json:"user_hash_id"`
	ChangedAt time.Time `json:"changed_at"`
}



func (m *Models) NewPreviousDataModel() string {
	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	    %s
	type VARCHAR(500) NOT NULL CHECK(type in ('email', 'username', 'phone', 'clan_tag')),
	value VARCHAR(500) NOT NULL,
	user_hash_id VARCHAR(500) NOT NULL
)`, database.PREVIOUS_DATA_TABLE, NewBaseModels())
}