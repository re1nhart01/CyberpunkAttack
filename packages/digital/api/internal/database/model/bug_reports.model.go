package models

import (
	"fmt"
	"github.com/cyberpunkattack/database"
)

type BugReportModel struct {
	*BaseModel
	Name string `json:"name"`
	Email string `json:"email"`
	Title string `json:"title"`
	Category string `json:"category"`
	Description string `json:"description"`
}



func (m *Models) NewBugReportTable() string {
	return fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
	%s
	name VARCHAR(500),
	email VARCHAR(500) NOT NULL,
	title VARCHAR(500) NOT NULL,
	category VARCHAR(500) NOT NULL,
	description TEXT NOT NULL,
	%s
)
`, database.BUG_REPORT_TABLE, NewBaseModels(), database.EmailConstaint)
}