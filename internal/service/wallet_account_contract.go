package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
)

type WalletAccountContract interface {
	EnableWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error)
	DisableWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error)
	ViewBallanceOnWallet(ctx context.Context, token string) (*presentation.WalletDataResponse, error)
}
