package routers

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	_ = repositories.NewRepository(se.DB)

}
