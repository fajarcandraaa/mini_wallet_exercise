package routers

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
	"github.com/fajarcandraaa/mini_wallet_exercise/usecase"
)

func walletRouter(p *PathPrefix, s *service.Service) {
	var (
		walletUseCase = usecase.NewWalletUseCase(s)
	)

	p.V1.HandleFunc("/init", walletUseCase.InitializeAccountWallet).Methods("POST")
}
