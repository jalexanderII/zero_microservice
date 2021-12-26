package server

import (
	"github.com/hashicorp/go-hclog"
	applicationDB "github.com/jalexanderII/zero_microservice/backend/services/application/database"
	applicationPB "github.com/jalexanderII/zero_microservice/gen/application"
)

type applicationServer struct {
	applicationPB.UnimplementedApplicationServer
	DB *applicationDB.ApplicationDB
	l  hclog.Logger
}

func NewApplicationServer(db *applicationDB.ApplicationDB, l hclog.Logger) *applicationServer {
	return &applicationServer{DB: db, l: l}
}
