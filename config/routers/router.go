package routers

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
)

func (se *Serve) initializeRoutes() {
	p := RouterConfigPrefix(se)            // set grouping prefix
	r := repositories.NewRepository(se.DB) //initiate repository
	s := service.NewService(r)             //initiate service

	// //initiate endpoint
	walletRouter(p, s)

}
