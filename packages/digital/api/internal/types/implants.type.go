package types

const (
	ImplantType           = "implant"
	ImplantTypeNeural     = "neural_system"
	ImplantTypeBody       = "body_system"
	ImplantTypeCybervirus = "cybervirus"
	ImplantTypeSkeleton   = "skeleton_system"
	ImplantTypeBlood      = "blood_system"
	ImplantTypeFight      = "fight_system"
	ImplantTypeOther      = "other_system"
)

type ParsedImplantsCardType struct {
	CardList []ActionCardCommonType `json:"as_list"`
	Implants struct {
		NeuralSystem   []ActionCardCommonType `json:"neural_system"`
		BodySystem     []ActionCardCommonType `json:"body_system"`
		Cybervirus     []ActionCardCommonType `json:"cybervirus"`
		SkeletonSystem []ActionCardCommonType `json:"skeleton_system"`
		BloodSystem    []ActionCardCommonType `json:"blood_system"`
		FightSystem    []ActionCardCommonType `json:"fight_system"`
		OtherSystem    []ActionCardCommonType `json:"other_system"`
	}
}

// TODO actualize
type ImplantCardCommonType struct {
	Id           int                      `json:"id"`
	Type         string                   `json:"type"`
	ImplantType  string                   `json:"implant_type"`
	Name         string                   `json:"name"`
	Spec         string                   `json:"spec"`
	Description  string                   `json:"description"`
	Illustration string                   `json:"illustration"`
	Target       string                   `json:"target"`
	Payload      []ImplantCardPayloadType `json:"payload"`
	Addon        string                   `json:"addon"`
}

type ImplantCardPayloadType struct {
}
