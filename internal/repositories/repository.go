package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Wallet WalletRepositoryContract
}

func NewWallet(db *gorm.DB) WalletRepositoryContract {
	return NewWalletRepository(db)
}

// NewRepository to setting repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Wallet: NewWallet(db),
	}
}
