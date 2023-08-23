package repositories

import (
	"context"
	"time"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/jinzhu/gorm"
)

type WalletAccountRepository struct {
	db *gorm.DB
}

func NewWalletAccountRepository(db *gorm.DB) *WalletAccountRepository {
	return &WalletAccountRepository{
		db: db,
	}
}

var _ WalletAccountRepositoryContract = &WalletAccountRepository{}

// UpdateStatus implements WalletAccountRepositoryContract.
func (r *WalletAccountRepository) UpdateStatus(ctx context.Context, status entity.WalletStatus, custromerXid string) (*entity.WalletAccount, error) {
	var (
		exist  bool
		result entity.WalletAccount
		t      = time.Now()
	)

	err := r.db.First(&result, "customer_xid = ? AND wallet_status = ?", custromerXid, status).Error
	if err != nil {
		exist = false
	} else {
		exist = true
	}

	if exist {
		return nil, entity.ErrWalletAlreadyExist
	}

	err = r.db.Model(&result).
		Where("customer_xid = ?", custromerXid).
		Updates(entity.WalletAccount{
			Status:    &status,
			EnabledAt: &t,
		}).Error
	if err != nil {
		return nil, err
	}

	err = r.db.First(&result, "customer_xid = ?", custromerXid).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetBalanceByCustomerXID implements WalletAccountRepositoryContract.
func (r *WalletAccountRepository) GetBalanceByCustomerXID(ctx context.Context, customerXID string) (*entity.WalletAccount, error) {
	var (
		walletAccount entity.WalletAccount
	)

	err := r.db.First(&walletAccount, "customer_xid = ? AND wallet_status = ?", customerXID, "enabled").Error
	if err != nil {
		return nil, entity.ErrWalletNotExist
	}

	return &walletAccount, nil
}
