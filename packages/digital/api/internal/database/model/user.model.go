package models

import (
	"fmt"

	"github.com/cyberpunkattack/database"
)

type PublicUserModel struct {
	*BaseModel
	UserHash string `json:"user_hash"` //not null unique
	Username string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role string `json:"role"`
	FullName string `json:"full_name"`
	Country string `json:"country"`
	City string `json:"city"`
	ClanTag *int `json:"clan_tag"`
	Description string `json:"description"`
}

type UserModel struct {
	*BaseModel
	UserHash string `json:"user_hash"` //not null unique
	Username string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role string `json:"role"`
	FullName string `json:"full_name"`
	Temporary bool `json:"temporary"`
	Country string `json:"country"`
	City string `json:"city"`
	ClanTag *int `json:"clan_tag"`
	Description string `json:"description"`
	Active bool `json:"active"`
	Password string `json:"password"`
}

func (m *Models) NewUserModel() string {
	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	    %s
		user_hash VARCHAR(500) NOT NULL UNIQUE,
		username VARCHAR(500) NOT NULL,
		email VARCHAR(100) UNIQUE,
		phone VARCHAR(500),
		full_name VARCHAR(500),
		country VARCHAR(500) DEFAULT 'World',
		role VARCHAR(500) NOT NULL DEFAULT 'User' CHECK(role in ('User', 'Admin')),
		city VARCHAR(500),
		clan_tag INT REFERENCES clans(id) ON DELETE SET NULL,
		description VARCHAR(2000),
		temporary BOOLEAN DEFAULT true,
		active BOOLEAN DEFAULT true,
		password VARCHAR(500) NOT NULL,
		%s,
		CONSTRAINT email_or_phone_check CHECK (
        email IS NOT NULL OR phone IS NOT NULL
    	)
);
CREATE UNIQUE INDEX IF NOT EXISTS unique_email ON users(email) WHERE email IS NOT NULL;
`, database.USERS_TABLE, NewBaseModels(), database.EmailConstaint)
}
