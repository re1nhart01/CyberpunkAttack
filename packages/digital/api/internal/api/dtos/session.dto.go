package dtos

import "github.com/cyberpunkattack/dto"

var CreateSessionDto = &dto.FieldsMapping{
	"invites": &dto.FieldDto{
		Type:      "ARRAY",
		Required:  true,
		Min:       0,
		Max:       8,
		Name:      "invites",
		MaxLength: 8,
		MinLength: 0,
	},
	"password": &dto.FieldDto{
		Name:     "password",
		Required: false,
		Type:     "STRING",
		Min:      2,
		Max:      100,
	},
	"name": &dto.FieldDto{
		Name:     "name",
		Required: false,
		Type:     "STRING",
		Min:      2,
		Max:      100,
	},
}
