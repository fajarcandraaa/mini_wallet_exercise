package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
)

type WalletAccountRepositoryContract interface {
	UpdateStatus(ctx context.Context, status entity.WalletStatus, custromerXid string) (*entity.WalletAccount, error)
}