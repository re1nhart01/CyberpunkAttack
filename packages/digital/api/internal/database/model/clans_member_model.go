package models

import (
	"fmt"

	"github.com/cyberpunkattack/database"
)

type ClansModelMember struct {
	*BaseModel
	ClanId     int    `json:"clan_id"`
	MemberHash string `json:"member_hash"`
	IsOwner    bool   `json:"is_owner"`
}

func (m *Models) NewClansMemberTable() string {
	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	%s
	clan_id INT REFERENCES clans(id) ON DELETE CASCADE,
	member_hash VARCHAR(500) REFERENCES users(user_hash) ON DELETE CASCADE,
	is_owner BOOL NOT NULL DEFAULT false
);
	CREATE UNIQUE INDEX IF NOT EXISTS unique_owner_per_clan
	ON clan_member (clan_id)
	WHERE (is_owner = true);
`, database.CLANS_MEMBER_TABLE, NewBaseModels())
}

// CONSTRAINT unique_owner_per_clan UNIQUE (clan_id, member_hash) WHERE (is_owner = true)
