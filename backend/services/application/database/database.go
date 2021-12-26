package applicationDB

import (
	"context"
	"database/sql"
	"fmt"
	"time"

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

// NewDBContext returns a new Context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*config.Performance/100)
}

func ConnectToTestDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=%s user=postgres sslmode=disable", config.TESTDBNAME, config.Pass))
	if err != nil {
		panic(err)
	}
	return db, err
}
