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

func TestTopUpVirtualMoney(t *testing.T) {
	db, rds, err := testConfig(t)
	require.NoError(t, err)
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{}, &entity.WalletTransaction{})

	r := repositories.NewRepository(db)
	walletTransactionService := service.NewWalletTransactionsService(r, rds)

	key, err := seeder.SeedEnabledWalletAccountFaker(db, rds)
	require.NoError(t, err)

	t.Run("feature view topup virtual money : if data is valid, expected no error", func(t *testing.T) {
		var (
			ctx    = context.Background()
			token  = faker.FakeToken
			amount = 15000
			reffID = faker.FakeReffID1
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletTransactionService.TopUpVirtualMoney(ctx, amount, reffID, tokenString)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("feature view topup virtual money : if data is duplicate, expected error", func(t *testing.T) {
		var (
			ctx    = context.Background()
			token  = faker.FakeToken
			amount = 15000
			reffID = faker.FakeReffID1
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		rsp, err := walletTransactionService.TopUpVirtualMoney(ctx, amount, reffID, tokenString)
		require.Error(t, err)
		assert.Nil(t, rsp)
	})

	_, err = rds.Del(context.Background(), key).Result()
	require.NoError(t, err)
}
