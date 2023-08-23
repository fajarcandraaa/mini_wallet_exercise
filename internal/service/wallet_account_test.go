package service_test

import (
	"context"
	"testing"

	"github.com/fajarcandraaa/mini_wallet_exercise/helpers"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
	"github.com/fajarcandraaa/mini_wallet_exercise/test/faker"
	"github.com/fajarcandraaa/mini_wallet_exercise/test/seeder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnabledAccountWallet(t *testing.T) {
	db, rds, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{})

	r := repositories.NewRepository(db)
	walletAccountService := service.NewWalletAccountService(r, rds)

	key, err := seeder.SeedWalletAccountFaker(db, rds)
	require.NoError(t, err)

	t.Run("feature enable wallet : if token is valid and exist, expected no error", func(t *testing.T) {
		var (
			ctx   = context.Background()
			token = faker.FakeToken
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletAccountService.EnableWallet(ctx, tokenString)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("feature enable wallet : if token is no valid or not exist, expected error", func(t *testing.T) {
		var (
			ctx   = context.Background()
			token = "Token 1234"
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletAccountService.EnableWallet(ctx, tokenString)
		require.Error(t, err)
		assert.Nil(t, res)
	})

	_, err = rds.Del(context.Background(), key).Result()
	require.NoError(t, err)
}

func TestViewWalletBallance(t *testing.T) {
	db, rds, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{})

	r := repositories.NewRepository(db)
	walletAccountService := service.NewWalletAccountService(r, rds)

	key, err := seeder.SeedEnabledWalletAccountFaker(db, rds)
	require.NoError(t, err)

	t.Run("feature view wallet ballance : if token is valid and exist, expected no error", func(t *testing.T) {
		var (
			ctx   = context.Background()
			token = faker.FakeToken
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletAccountService.ViewBallanceOnWallet(ctx, tokenString)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("feature view wallet ballance : if token is no valid or not exist, expected error", func(t *testing.T) {
		var (
			ctx   = context.Background()
			token = "Token i81723yjb-213jkgkweg"
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletAccountService.ViewBallanceOnWallet(ctx, tokenString)
		require.Error(t, err)
		assert.Nil(t, res)
	})

	_, err = rds.Del(context.Background(), key).Result()
	require.NoError(t, err)
}
