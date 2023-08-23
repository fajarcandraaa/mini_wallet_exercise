package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
)

type WalletTransactionsRepositoryContract interface {
	AddBalance(ctx context.Context, p entity.WalletTransaction) (*entity.WalletTransaction, error)
	ReduceBalance(ctx context.Context, p entity.WalletTransaction) error
}
