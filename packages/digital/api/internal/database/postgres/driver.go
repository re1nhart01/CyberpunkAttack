package postgres

import (
  "fmt"
  "github.com/cyberpunkattack/environment"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type PostgresDb struct {
	instance *gorm.DB
}

var pgDbInstance *PostgresDb = nil

func New() *PostgresDb {

  POSTGRES_PORT := environment.GEnv().Get("POSTGRES_PORT")
  POSTGRES_HOST := environment.GEnv().Get("POSTGRES_HOST")
  POSTGRES_PASSWORD := environment.GEnv().Get("POSTGRES_PASSWORD")
  POSTGRES_USER := environment.GEnv().Get("POSTGRES_USER")
  POSTGRES_DB := environment.GEnv().Get("DB_NAME")

  db, err := gorm.Open(postgres.New(postgres.Config{
    DSN: fmt.Sprintf(
      "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
      POSTGRES_HOST,
      POSTGRES_USER,
      POSTGRES_PASSWORD,
      POSTGRES_DB,
      POSTGRES_PORT,
      ),
    PreferSimpleProtocol: true, // disables implicit prepared statement usage
  }), &gorm.Config{})

  if err != nil {
    panic("CAN NOT CONNECT TO POSTGRES DB" + err.Error())
  }

  pgDbInstance = &PostgresDb{
    instance: db,
  }

  return pgDbInstance
}


func DB() *PostgresDb {
  return pgDbInstance
}

func (db *PostgresDb) Get() *gorm.DB {
  return db.instance
}

func (db *PostgresDb) Clear() {
  pgDbInstance = nil
}