package main

import (
	"fmt"

	nsg "github.com/cyberpunkattack/pkg/no-stupid-gen"
)

func main() {
	g := &nsg.NoStupidGen{
		"name": {
			Name:        "name",
			Datatype:    nsg.Datatypes.VARCHAR,
			LenIfExists: 100,
			Constraints: []string{nsg.Constraints.UNIQUE, nsg.Constraints.NOT_NULL},
			OnUpdate:    nsg.Actions.SET_NULL,
		},
		"password": {
			Name:        "password",
			Datatype:    nsg.Datatypes.VARCHAR,
			LenIfExists: 100,
			Constraints: []string{nsg.Constraints.UNIQUE, nsg.Constraints.NOT_NULL, nsg.ConstraintDefault(nsg.Datatypes.VARCHAR, "aboba")},
			OnUpdate:    nsg.Actions.SET_NULL,
		},
	}
	noGen := nsg.New(g)
	fmt.Println(noGen.CreateTable("users", &nsg.NoStupidConfig{
		IfNotExists: true,
	}))
}
