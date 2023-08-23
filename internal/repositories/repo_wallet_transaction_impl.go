package repositories

import (
	"context"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/jinzhu/gorm"
)

type WalletTransactionsRepository struct {
	db *gorm.DB
}

func NewWalletTransactionsRepository(db *gorm.DB) *WalletTransactionsRepository {
	return &WalletTransactionsRepository{
		db: db,
	}
}

var _ WalletTransactionsRepositoryContract = &WalletTransactionsRepository{}

// AddBalance implements WalletTransactionsRepositoryContract.
func (r *WalletTransactionsRepository) AddBalance(ctx context.Context, p entity.WalletTransaction) (*entity.WalletTransaction, error) {
	var (
		result      entity.WalletTransaction
		exist       = false
		reffID      = p.ReferenceID
		querySearch = `SELECT wallet_trx_id FROM wallet_transactions WHERE reference_id = $1`
		queryTrx    = `
			INSERT INTO wallet_transactions
			(wallet_trx_id, wallet_id, wallet_trx_type, wallet_ballance_trx, deposited_by, deposited_at, reference_id)
			VALUES ($1,$2,$3,$4,$5,$6,$7);
		`
		queryUpdate = `
			UPDATE wallets SET wallet_ballance = wallet_ballance + $1 WHERE customer_xid = $2;
		`
	)
	arg := []interface{}{
		&p.WalletTrxID,
		&p.WalletID,
		&p.WalletTrxType,
		&p.WalletBallanceTrx,
		&p.DepositedBy,
		&p.DepositedAt,
		&p.ReferenceID,
	}
	tx := r.db.Begin()
	// _, err := tx.Raw("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE").Rows()
	err := tx.Exec("SET TRANSACTION ISOLATION LEVEL SERIALIZABLE").Error
	if err != nil {
		return nil, err
	}

	ex := tx.Exec(querySearch, reffID)
	if ex.RowsAffected > 0 {
		exist = true
	}

	if exist {
		tx.Rollback()
		return nil, entity.ErrWalletAlreadyExist
	}

	err = tx.Exec(queryTrx, arg...).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	argUpdate := []interface{}{
		&p.WalletBallanceTrx,
		&p.DepositedBy,
	}

	err = tx.Exec(queryUpdate, argUpdate...).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return &result, nil

}

// ReduceBalance implements WalletTransactionsRepositoryContract.
func (*WalletTransactionsRepository) ReduceBalance(ctx context.Context, p entity.WalletTransaction) error {
	panic("unimplemented")
}
