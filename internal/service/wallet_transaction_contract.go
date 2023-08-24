package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
)

type WalletTransactionContract interface {
	TopUpVirtualMoney(ctx context.Context, amount int, reffID, token string) (*presentation.DepositOrWithdrawlResponse, error)
	UseVirtualMoney(ctx context.Context, amount int, reffID, token string) (*presentation.DepositOrWithdrawlResponse, error)
}
