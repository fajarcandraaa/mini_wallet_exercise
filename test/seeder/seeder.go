package seeder

import (
	"github.com/fajarcandraaa/mini_wallet_exercise/test/faker"
	"github.com/jinzhu/gorm"
)

func SeedWalletAccountFaker(db *gorm.DB) error {
	fakeWalletAccount := faker.FakeWalletAccount()
	err := db.Create(&fakeWalletAccount).Error
	if err != nil {
		return err
	}
	return nil
}
