package applicationDB

import (
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	config "github.com/jalexanderII/zero_microservice"
	"github.com/jalexanderII/zero_microservice/backend/services/application/database/genDB"
	_ "github.com/lib/pq"
)

type ApplicationDB struct {
	*genDB.Queries
	DB *sql.DB
}

func NewApplicationDB(db *sql.DB) *ApplicationDB {
	return &ApplicationDB{DB: db, Queries: genDB.New(db)}
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=%s user=postgres sslmode=disable", config.ApplicationDBNAME, config.Pass))
	if err != nil {
		panic(err)
	}
	return db, err
}
