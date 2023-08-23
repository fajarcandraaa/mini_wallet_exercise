package dto

import (
	"time"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/google/uuid"
)

func WalletTrxToResponse(p entity.WalletTransaction) *presentation.DepositResponse {
	detail := presentation.TrxDetailResponse{
		ID:          p.WalletTrxID,
		DepositedBy: p.DepositedBy,
		Status:      "success",
		DepositedAt: p.DepositedAt,
		Amount:      p.WalletBallanceTrx,
		ReffID:      p.ReferenceID,
	}

	resp := presentation.DepositResponse{
		Deposit: detail,
	}

	return &resp
}

func AddBalanceRequest(amount int, reffID string) presentation.AddBalanceRequest {
	resp := presentation.AddBalanceRequest{
		Amount:  amount,
		ReffID:  reffID,
		TrxType: entity.Deposit,
	}

	return resp
}

func AddBalanceRequestToDatabase(p presentation.AddBalanceRequest, d presentation.CustomerDataByTokenResponse) *entity.WalletTransaction {
	t := time.Now()
	res := entity.WalletTransaction{
		WalletTrxID:       uuid.NewString(),
		WalletID:          d.WalletID,
		WalletTrxType:     &p.TrxType,
		WalletBallanceTrx: p.Amount,
		DepositedBy:       d.CustomerID,
		DepositedAt:       &t,
		ReferenceID:       p.ReffID,
	}

	return &res
}

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

func WalletAccountToResponse(p entity.WalletAccount) presentation.WalletDataResponse {
	res := presentation.WalletDetailDataResponse{
		ID:        p.AccountID,
		OwnedBy:   p.CustomerXid,
		Status:    string(*p.Status),
		EnabledAt: *p.EnabledAt,
		Balance:   p.WalletBallance,
	}

	result := presentation.WalletDataResponse{
		Wallet: res,
	}

	return result
}

func TokenToResponse(t string) presentation.InitiateWalletAccountResponse {
	res := presentation.InitiateWalletAccountResponse{
		Token: t,
	}

	return res
}
