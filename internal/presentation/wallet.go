package presentation

import (
	"time"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
)

// Request
type (
	InitiateWalletAccountRequest struct {
		CustomerXid string `json:"customer_xid"`
	}

	NewWalletAccountRequest struct {
		AccountID      string              `json:"account_id"`
		WalletID       string              `json:"wallet_id"`
		CustomerXid    string              `json:"customer_xid"`
		WalletStatus   entity.WalletStatus `json:"wallet_status"`
		WalletBallance int                 `json:"wallet_ballance"`
	}
)

// Response
type (
	InitiateWalletAccountResponse struct {
		Token string `json:"token"`
	}

	WalletDataResponse struct {
		ID        string `json:"id"`
		OwnedBy   string `json:"owned_by"`
		Status    string `json:"status"`
		EnabledAt string `json:"enabled_at"`
		Balance   int    `json:"balance"`
	}

	WalletTransactionDetailResponse struct {
		ID                string                `json:"id"`
		WalletID          string                `json:"wallet_id"`
		WalletTrxType     *entity.WalletTrxType `json:"wallet_trx_type"`
		WalletBallanceTrx int                   `json:"wallet_ballance_trx"`
		DepositedBy       string                `json:"deposited_by,omitempty"`
		DepositedAt       *time.Time            `json:"deposited_at,omitempty"`
		WithdrawnBy       string                `json:"withdrawn_by,omitempty"`
		WithdrawnAt       *time.Time            `json:"withdrawn_at,omitempty"`
		CreatedAt         time.Time             `json:"created_at,omitempty"`
		UpdatedAt         time.Time             `json:"updated_at,omitempty"`
	}

	WalletTransactionListResponse struct {
		Transactions []WalletTransactionDetailResponse `json:"transactions"`
	}
)
