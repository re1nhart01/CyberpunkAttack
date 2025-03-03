package types

type PayloadType string

const (
	SpecCorporateBoss = "roles_corporates_corporate-boss"
	SpecCorporate     = "roles_corporates_corporate"
	SpecCyberstray    = "roles_cyberstray_cyberstray"
	SpecLaidFlicker   = "roles_cyberpunks_laid-flicker"
	SpecCyberpunk     = "roles_cyberpunks_cyberpunk"
)

const (
	EffectAdditionalImplants   PayloadType = "additional_implants"
	EffectAdditionalCardOnDraw PayloadType = "additional_card_on_draw"
	EffectAdditionalMaxHealth  PayloadType = "additional_max_health"
)

type RoleCardCommonType struct {
	Id           int                    `json:"id"`
	Type         string                 `json:"type"`
	RolesFaction string                 `json:"roles_faction"`
	Name         string                 `json:"name"`
	Spec         string                 `json:"spec"`
	Description  string                 `json:"description"`
	Illustration string                 `json:"illustration"`
	WinCondition *string                `json:"winCondition"`
	Payload      []*RoleCardPayloadType `json:"payload"`
}

type RoleCardPayloadType struct {
	EffectName     PayloadType `json:"effect_name"`
	HowMuch        int         `json:"howMuch"`
	SkipViruses    bool        `json:"skip_viruses,omitempty"`
	ConditionToGet string      `json:"conditionToGet"`
}
