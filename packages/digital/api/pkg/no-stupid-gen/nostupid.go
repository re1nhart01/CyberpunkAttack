package nsg

import (
	"fmt"
	"slices"
	"strings"
)

type NoStupidGen map[string]*NoStupidField
type NoStupidField struct {
	Name        string
	Datatype    string
	Constraints []string
	OnDelete    string
	OnUpdate    string
	LenIfExists int
	Indexes     []string
}

type NoStupidConfig struct {
	GenStringOnly bool
	IfNotExists   bool
	AddIDSerial   bool
	AddUpdatedAt  bool
	AddDeletedAt  bool
	AddCreateddAt bool
}

func ConstraintDefault(d, v string) string {
	addV := slices.Contains([]string{Datatypes.VARCHAR, Datatypes.TEXT, Datatypes.CHAR}, d)
	c := ternary(addV, "'", "")
	return fmt.Sprintf("DEFAULT %s%s%s", c, v, c)
}

func ternary[T comparable](v bool, v1 T, v2 T) T {
	if v {
		return v1
	} else {
		return v2
	}
}

func generateOnUpdate(v1 string) string {
	if len(v1) < 1 {
		return ""
	}
	return fmt.Sprintf("%s %s ", Constraints.ON_UPDATE, v1)
}

func generateOnDelete(v1 string) string {
	if len(v1) < 1 {
		return ""
	}
	return fmt.Sprintf("%s %s ", Constraints.ON_DELETE, v1)
}

func generateDatatype(v1 string, len int) string {
	hasLen := slices.Contains([]string{
		Datatypes.VARCHAR,
		Datatypes.CHAR,
		Datatypes.BINARY,
		Datatypes.VARBINARY,
		Datatypes.DECIMAL,
		Datatypes.NUMERIC,
	}, v1)
	if !hasLen {
		return v1
	}
	return fmt.Sprintf("%s(%d)", v1, len)
}

func (g *NoStupidGen) CreateTable(name string, config *NoStupidConfig) string {
	ifNotExistsString := ternary[string](config.IfNotExists, "IF NOT EXISTS", "")
	tableImplementation := "\n"

	counter := len(*g)
	current := 0

	for _, v := range *g {
		field := *v
		fieldText := ""
		fieldText += field.Name + " "
		fieldText += generateDatatype(field.Datatype, field.LenIfExists) + " "
		fieldText += strings.Join(field.Constraints, " ") + " "
		fieldText += generateOnDelete(field.OnDelete)
		fieldText += generateOnUpdate(field.OnUpdate)
		current++
		if current != counter {
			fieldText += ","
		}
		tableImplementation += fieldText + "\n"
	}

	resultTableString := fmt.Sprintf("CREATE TABLE %s %s (%s)", ifNotExistsString, name, tableImplementation)
	return resultTableString
}

func New(gen *NoStupidGen) *NoStupidGen {
	return gen
}
