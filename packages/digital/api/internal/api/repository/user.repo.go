package repository

import (
	"fmt"

	"github.com/cyberpunkattack/api/base"
	"github.com/cyberpunkattack/database"
	models "github.com/cyberpunkattack/database/model"
	"github.com/cyberpunkattack/database/postgres"
)

type UserRepository struct {
	*base.Repository
	injections *UserInjections
}

type UserInjections struct {
	Clans ClansRepository
}

func (user *UserRepository) GetUserByField(field string, value any) (*models.UserModel, error) {
	userModel := models.UserModel{}
	payload := postgres.DB().Get().Table(database.USERS_TABLE).Where(fmt.Sprintf("%s = ?", field), value).First(&userModel)
	if payload.Error != nil {
		return &userModel, payload.Error
	}

	return &userModel, nil
}

func (user *UserRepository) GetUserByEmail(email string) (*models.UserModel, bool) {
	userModel := models.UserModel{}
	payload := postgres.DB().Get().Table(database.USERS_TABLE).Where("email = ?", email).First(&userModel)
	if payload.Error != nil {
		return &userModel, false
	}

	return &userModel, true
}

func (user *UserRepository) GetFriendsListByUserHash(userHash string) ([]models.FriendModel, error) {
	result := []models.FriendModel{}

	//TODO: implement select list

	return result, nil
}

func (user *UserRepository) GetUserByUserHash(userHash string, isMe bool) (*models.CompoundUserProfile, error) {
	result := &models.CompoundUserProfile{}

	if userData, err := user.GetUserByField("user_hash", userHash); err != nil {
		return result, err
	} else {
		result.UserData = models.ToObfuscateUser(*userData)
	}

	if clanData, isExists, _ := user.injections.Clans.GetClanByUserHash(userHash); isExists {
		result.CurrentClan = clanData
	}

	if friendList, err := user.GetFriendsListByUserHash(userHash); err != nil {
		return result, err
	} else {
		result.Friends = friendList
	}

	return result, nil
}

func NewUserRepository(injections *UserInjections) *UserRepository {
	return &UserRepository{
		Repository: &base.Repository{
			TableName: "users",
		},
		injections: injections,
	}
}
