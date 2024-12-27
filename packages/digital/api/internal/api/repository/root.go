package repository

type Injectable struct {
	Auth  AuthRepository
	User  UserRepository
	App   AppRepository
	Clans ClansRepository
}

type UserCredentials struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpireIn     int64  `json:"expireIn"`
}
