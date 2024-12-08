package dtos

import (
    "github.com/cyberpunkattack/dto"
    "regexp"
)

var InitialSignUpDto = &dto.FieldsMapping{
	"email": &dto.FieldDto{
		Name: "email",
		Required: false,
		Type: "STRING",
		Min: 5,
		Max: 50,
		RegexpValidation: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	},
	"phone": &dto.FieldDto{
		Name: "phone",
		Required: false,
		Type: "STRING",
		Min: 10,
		Max: 20,
	},
	"username": &dto.FieldDto{
		Name: "username",
		Required: true,
		Type: "STRING",
		Min: 3,
		Max: 100,
	},
	"fullName": &dto.FieldDto{
		Name: "fullName",
		Required: true,
		Type: "STRING",
		Min: 5,
		Max: 100,
	},
	"password": &dto.FieldDto{
		Name: "password",
		Required: true,
		Type: "STRING",
		Min: 2,
		Max: 300,
	},
}


var ValidateEmailDto = &dto.FieldsMapping{
	"code": &dto.FieldDto{
		Name: "code",
		Required: true,
		Type: "STRING",
		Min: 10,
		Max: 100,
	},
	"email": &dto.FieldDto{
		Name: "email",
		Required: true,
		Type: "STRING",
		Min: 5,
		Max: 50,
		RegexpValidation: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	},
}


var LogInDto = &dto.FieldsMapping{
	"email": &dto.FieldDto{
		Name: "email",
		Required: true,
		Type: "STRING",
		Min: 5,
		Max: 50,
		RegexpValidation: regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
	},
	"password": &dto.FieldDto{
		Name: "password",
		Required: true,
		Type: "STRING",
		Min: 2,
		Max: 300,
	},
}



var RefreshTokenDto = &dto.FieldsMapping{
	"refreshToken": &dto.FieldDto{
		Name: "refreshToken",
		Required: true,
		Type: "STRING",
		Min: 2,
		Max: 300,
	},
	"grantType": &dto.FieldDto{
		Name: "grantType",
		Required: true,
		Type: "STRING",
		Min: 2,
		Max: 15,
	},
}