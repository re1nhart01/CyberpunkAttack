package models

func ToObfuscateUser(user UserModel) *PublicUserModel {
	return &PublicUserModel{
		BaseModel:   user.BaseModel,
		UserHash:    user.UserHash,
		Username:    user.Username,
		Email:       user.Email,
		Phone:       user.Phone,
		Role:        user.Role,
		FullName:    user.FullName,
		Country:     user.Country,
		City:        user.City,
		ClanTag:     user.ClanTag,
		Description: user.Description,
	}
}
