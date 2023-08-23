package service

import "github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"

type Service struct {
	WalletService WalletServiceContract
}
func ServiceWallet(repo *repositories.Repository) WalletServiceContract {
	return NewWalletService(repo)
}
func NewService(repo *repositories.Repository) *Service {
	return &Service{
		WalletService: ServiceWallet(repo),
	}
}
