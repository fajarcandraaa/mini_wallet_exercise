package service

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
)

type WalletServiceContract interface {
	CreateAccount(ctx context.Context, payload presentation.InitiateWalletAccountRequest) (*presentation.InitiateWalletAccountResponse, error)
}