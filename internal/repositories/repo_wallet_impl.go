package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/helpers"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/jinzhu/gorm"
)

type WalletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *WalletRepository {
	return &WalletRepository{
		db: db,
	}
}

var _ WalletRepositoryContract = &WalletRepository{}

// GetDataCustomerByToken implements WalletRepositoryContract.
func (r *WalletRepository) GetDataCustomerByToken(ctx context.Context, customerXid string) (*presentation.CustomerDataByTokenResponse, error) {
	var (
		result presentation.CustomerDataByTokenResponse
		model  entity.Wallet
	)

	err := r.db.Model(&model).
		Select("wallets.wallet_id, wallets.customer_xid, wallet_accounts.account_id").
		Joins("JOIN wallet_accounts on wallets.customer_xid = wallet_accounts.customer_xid").
		Where("wallets.customer_xid = ? ", customerXid).
		Scan(&result).Error
	if err != nil {
		return nil, entity.ErrWalletNotExist
	}

	result.CustomerID = customerXid

	return &result, nil
}

// StoreNewWallet implements WalletRepositoryContract.
func (w *WalletRepository) StoreNewWallet(ctx context.Context, payload presentation.NewWalletAccountRequest) (*string, error) {
	var (
		queryWallet = `
			INSERT INTO wallets (wallet_id, customer_xid, wallet_ballance) VALUES ($1, $2, $3);
		`
		queryWalletAccount = `
			INSERT INTO wallet_accounts (account_id, customer_xid, wallet_status, wallet_ballance) VALUES ($1, $2, $3, $4)
		`
	)

	argWallet := []interface{}{
		&payload.WalletID,
		&payload.CustomerXid,
		0,
	}

	argAccount := []interface{}{
		&payload.AccountID,
		&payload.CustomerXid,
		&payload.WalletStatus,
		0,
	}
	tx := w.db.Begin()
	_, err := tx.Raw("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE").Rows()
	if err != nil {
		return nil, err
	}

	err = insertWallet(tx, queryWallet, argWallet)
	if err != nil {
		return nil, err
	}

	err = insertWalletAccount(tx, queryWalletAccount, argAccount)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateHexadecimalStringTokent()
	if err != nil {
		return nil, err
	}

	return token, nil
}

func insertWallet(tx *gorm.DB, q string, arg []interface{}) error {
	err := tx.Exec(q, arg...).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func insertWalletAccount(tx *gorm.DB, q string, arg []interface{}) error {
	err := tx.Exec(q, arg...).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
