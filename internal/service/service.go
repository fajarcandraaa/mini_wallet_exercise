package service

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	WalletService WalletServiceContract
	WalletAccount WalletAccountContract
}

func ServiceWallet(repo *repositories.Repository, rds *redis.Client) WalletServiceContract {
	return NewWalletService(repo, rds)
}

func ServiceWalletAccount(repo *repositories.Repository, rds *redis.Client) WalletAccountContract {
	return NewWalletAccountService(repo, rds)
}

func NewService(repo *repositories.Repository, rds *redis.Client) *Service {
	return &Service{
		WalletService: ServiceWallet(repo, rds),
		WalletAccount: ServiceWalletAccount(repo, rds),
	}
}
