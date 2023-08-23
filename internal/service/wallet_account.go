package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/helpers"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/dto"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	"github.com/go-redis/redis/v8"
)

type walletAccount struct {
	repo *repositories.Repository
	rds  *redis.Client
}

func NewWalletAccountService(repo *repositories.Repository, rds *redis.Client) *walletAccount {
	return &walletAccount{
		repo: repo,
		rds:  rds,
	}
}

var _ WalletAccountContract = &walletAccount{}

// EnableWallet implements WalletAccountContract.
func (s *walletAccount) EnableWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error) {
	var (
		tokenKey string
	)
	keys, err := s.rds.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	for _, k := range keys {
		v, err := s.rds.Get(ctx, k).Result()
		if err != nil {
			return nil, err
		}
		if v == token {
			tokenKey = k
			break
		}
	}

	if tokenKey == "" {
		return nil, entity.ErrPermissionNotAllowed
	}

	customerXid := helpers.GetCustomerXidFromToken(tokenKey)

	enabledWallet, err := s.repo.WalletAccount.UpdateStatus(ctx, entity.Enable, customerXid)
	if err != nil {
		return nil, err
	}

	rsp := dto.WalletAccountToResponse(*enabledWallet)

	return &rsp, nil

}
