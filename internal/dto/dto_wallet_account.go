package dto

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/google/uuid"
)

func CustomerXidToDatabase(p presentation.InitiateWalletAccountRequest) *presentation.NewWalletAccountRequest {
	res := &presentation.NewWalletAccountRequest{
		AccountID:      uuid.NewString(),
		WalletID:       uuid.NewString(),
		CustomerXid:    p.CustomerXid,
		WalletStatus:   entity.Disable,
		WalletBallance: 0,
	}

	return res
}

func TokenToResponse(t string) presentation.InitiateWalletAccountResponse {
	res := presentation.InitiateWalletAccountResponse{
		Token: t,
	}

	return res
}
