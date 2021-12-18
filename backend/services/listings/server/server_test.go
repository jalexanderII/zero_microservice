package server

import (
	"github.com/hashicorp/go-hclog"
	contentStore "github.com/jalexanderII/zero_microservice/backend/services/listings/store"
)

var L = hclog.Default()
var Store = contentStore.NewDiskImageStore("./Store/tmp", L)
