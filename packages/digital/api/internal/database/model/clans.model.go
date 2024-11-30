package models

import (
	"fmt"
	"github.com/cyberpunkattack/database"
)

type ClansModel struct {
	*BaseModel
	Name string `json:"name"`
	SmallName string `json:"small_name"`
	Description string `json:"description"`
	Lang string `json:"lang"`
	Region string `json:"region"`
}


func (m *Models) NewClansTable() string {
	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	%s
	name VARCHAR(500) NOT NULL UNIQUE,
	small_name VARCHAR(10) NOT NULL UNIQUE,
	description VARCHAR(2000),
	lang VARCHAR(4) NOT NULL DEFAULT 'en',
	region VARCHAR(500) NOT NULL
)
`, database.CLANS_TABLE, NewBaseModels())
}