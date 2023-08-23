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
		result entity.WalletAccount
		t      = time.Now()
	)
	err := r.db.Model(&result).
		Where("customer_xid = ?", custromerXid).
		Updates(entity.WalletAccount{
			Status:    &status,
			EnabledAt: &t,
		}).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Debug().First(&result, "customer_xid = ?", custromerXid).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
