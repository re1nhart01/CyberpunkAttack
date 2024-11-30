package postgres

import "fmt"

func CreateSpecificDatabase(name string) {
	curr := DB()
	curr.Get().Exec(fmt.Sprintf("CREATE DATABASE %s;", name))
	curr.Eliminate()
	defer func() {
		err := New()
		if err != nil {
			panic("CreateSpecificDatabase error")
		}
	}()
}

func (db *PostgresDb) Eliminate() {
	if db == nil {
		return
	}
	sqlDb, err := db.Get().DB()
	if err != nil {
		panic("err on Eliminate method")
	}
	_ = sqlDb.Close()
	defer db.Clear()
}


func (db *PostgresDb) CallManualSQL(query string) {
	db.Get().Exec(query)
}