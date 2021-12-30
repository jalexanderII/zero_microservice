package listingsDB

import (
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	config "github.com/jalexanderII/zero_microservice"
	"github.com/jalexanderII/zero_microservice/backend/services/listings/database/genDB"
	_ "github.com/lib/pq"
)

type ListingsDB struct {
	*genDB.Queries
	DB *sql.DB
}

func NewListingsDB(db *sql.DB) *ListingsDB {
	return &ListingsDB{DB: db, Queries: genDB.New(db)}
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=%s user=postgres sslmode=disable", config.DBNAME, config.Pass))
	if err != nil {
		panic(err)
	}
	return db, err
}
