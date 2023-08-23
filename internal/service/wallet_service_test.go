package service_test

import (
	"context"
	"testing"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsetNewWalletAccount(t *testing.T) {
	var (
		ctx = context.Background()
		dsn = "host=localhost user=postgres dbname=julo_backend_test sslmode=disable password=postgres port=5433"
	)

	db, err := gorm.Open("postgres", dsn)
	require.NoError(t, err)
	db.AutoMigrate(&entity.Wallet{}, &entity.WalletAccount{})
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{})

	r := repositories.NewRepository(db)
	walletService := service.NewWalletService(r)

	t.Run("if store data given valid data, expected no error", func(t *testing.T) {
		var (
			payload = presentation.InitiateWalletAccountRequest{
				CustomerXid: "ea0212d3-abd6-406f-8c67-868e814a2436",
			}
		)
		res, err := walletService.CreateAccount(ctx, payload)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("if store data given not valid data, expected  error", func(t *testing.T) {
		var (
			payload = presentation.InitiateWalletAccountRequest{
				CustomerXid: "123456",
			}
		)
		res, err := walletService.CreateAccount(ctx, payload)
		require.Error(t, err)
		assert.Nil(t, res)
	})

}
