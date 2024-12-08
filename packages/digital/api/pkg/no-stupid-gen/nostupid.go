package no_stupid_gen

import "fmt"

type NoStupidField struct {
	Name string
	Datatype string
	Constraints []string
	OnDelete string
	OnUpdate string
	LenIfExists int
	Indexes []string
}

type NoStupidConfig struct {

}

type NoStupidGen map[string]*NoStupidField


func ConstraintDefault(v string) string {
	return fmt.Sprintf("DEFAULT %s", v)
}


func (g *NoStupidGen) CreateTable(name string, config NoStupidConfig) string {
	return ``
}