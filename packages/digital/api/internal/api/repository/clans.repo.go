package repository

import (
	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/database"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/postgres"
)

/*

Name        string `json:"name"`
	SmallName   string `json:"small_name"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Region      string `json:"region"`
	ClanId      int    `json:"clan_id"`
	MemberHash  string `json:"member_hash"`
	IsOwner     bool   `json:"is_owner"`
*/

type ClansRepository struct {
	*base.Repository
}

func (clans *ClansRepository) GetClanByUserHash(userHash string) (*models.CompoundClanWithMember, bool, error) {
	result := &models.CompoundClanWithMember{}

	if payload := postgres.DB().Get().Table("? as cm", database.CLANS_MEMBER_TABLE).Select(`
		cm.clan_id, cm.member_hash, cm.is_owner,
		clans.name, clans.small_name, clans.description,
	 	clans.lang, clans.region`).
		Joins("inner join clans on clans.id = cm.clan_id").First(&result); payload.Error != nil {
		return result, false, payload.Error
	}

	return result, true, nil
}

func NewClansRepository() *ClansRepository {
	return &ClansRepository{
		Repository: &base.Repository{TableName: "clans"},
	}
}
