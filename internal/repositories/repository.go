package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Wallet        WalletRepositoryContract
	WalletAccount WalletAccountRepositoryContract
}

func NewWallet(db *gorm.DB) WalletRepositoryContract {
	return NewWalletRepository(db)
}

func NewWalletAccount(db *gorm.DB) WalletAccountRepositoryContract {
	return NewWalletAccountRepository(db)
}

// NewRepository to setting repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Wallet:        NewWallet(db),
		WalletAccount: NewWalletAccount(db),
	}
}
