package models

import (
	"fmt"

	"github.com/cyberpunkattack/database"
)

const (
	ACCEPTED_FRIEND = "friends"
	BLOCKED_FRIEND  = "blocked"
)

const STATUSES = "'" + ACCEPTED_FRIEND + "'" + "," + "'" + BLOCKED_FRIEND + "'"

type FriendModel struct {
	*BaseModel
	UserHash    string `json:"user_hash"` //not null unique
	Username    string `json:"username"`
	Email       string `json:"email"`
	ClanTag     *int   `json:"clan_tag"`
	Description string `json:"description"`
	FullName    string `json:"full_name"`
}

func (m *Models) NewFriendModel() string {
	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	%s
	user_id INT REFERENCES users(id) ON DELETE CASCADE,
	friend_id INT REFERENCES users(id) ON DELETE CASCADE,
	status VARCHAR(100) NOT NULL DEFAULT '%s' CHECK(status in (%s)),
	CONSTRAINT check_is_unordered CHECK(user_id < friend_id),
	CONSTRAINT unique_ids UNIQUE(user_id, friend_id)
)
`, database.FRIENDS_TABLE, NewBaseModels(), ACCEPTED_FRIEND, STATUSES)
}

//	user_hash VARCHAR(500) NOT NULL UNIQUE,
//	username VARCHAR(500) NOT NULL,
//	email VARCHAR(100) UNIQUE,
//	clan_tag INT REFERENCES clans(id) ON DELETE SET NULL,
//	description VARCHAR(2000),
//	full_name VARCHAR(500)
