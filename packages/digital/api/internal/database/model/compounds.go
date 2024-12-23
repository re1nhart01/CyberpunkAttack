package models

type CompoundUserProfile struct {
	UserData    PublicUserModel        `json:"user_data"`
	CurrentClan CompouldClanWithMember `json:"current_clan"`
	Friends     []FriendModel          `json:"friends"`
}

type CompouldClanWithMember struct {
	*BaseModel
	Name        string `json:"name"`
	SmallName   string `json:"small_name"`
	Description string `json:"description"`
	Lang        string `json:"lang"`
	Region      string `json:"region"`
	ClanId      int    `json:"clan_id"`
	MemberHash  string `json:"member_hash"`
	IsOwner     bool   `json:"is_owner"`
}
