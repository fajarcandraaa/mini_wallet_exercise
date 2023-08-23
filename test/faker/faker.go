package faker

import (
	"time"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
)

const (
	FakeToken                      = "Token 1d54be82cbad3ec463dcd0ceab26fb409d5e4a52"
	FakeTokenValue                 = "1d54be82cbad3ec463dcd0ceab26fb409d5e4a52"
	FakeAccountID                  = "607bf44c-102a-4f4b-89ce-4b75d67ce415"
	FakeCustomerXID                = "ea0212d3-abd6-406f-8c67-868e814a2436"
	FakeWalletID                   = "fc70c26a-5922-4f83-843f-61859f9cf55c"
	FakeWalletAccountStatusDisable = entity.Disable
	FakeWalletAccountStatusEnable  = entity.Enable
)

func FakeWalletAccount() *entity.WalletAccount {

	var (
		status entity.WalletStatus
		t      = time.Now()
	)
	status = FakeWalletAccountStatusDisable

	fakeWalletAccount := &entity.WalletAccount{
		AccountID:      FakeAccountID,
		CustomerXid:    FakeCustomerXID,
		Status:         &status,
		WalletBallance: 0,
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	return fakeWalletAccount
}

func FakeWalletAccountEnable() *entity.WalletAccount {

	var (
		status entity.WalletStatus
		t      = time.Now()
	)
	status = FakeWalletAccountStatusEnable

	fakeWalletAccount := &entity.WalletAccount{
		AccountID:      FakeAccountID,
		CustomerXid:    FakeCustomerXID,
		Status:         &status,
		WalletBallance: 0,
		EnabledAt:      &t,
		CreatedAt:      t,
		UpdatedAt:      t,
	}

	return fakeWalletAccount
}
