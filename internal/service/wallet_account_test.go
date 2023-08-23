package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/fajarcandraaa/mini_wallet_exercise/helpers"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
	"github.com/fajarcandraaa/mini_wallet_exercise/test/seeder"
	rds "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnabledAccountWallet(t *testing.T) {
	var (
		ctx = context.Background()
		dsn = "host=localhost user=postgres dbname=julo_backend_test sslmode=disable password=postgres port=5433"
	)

	db, err := gorm.Open("postgres", dsn) // initiate database for testing
	require.NoError(t, err)
	db.AutoMigrate(&entity.Wallet{}, &entity.WalletAccount{})
	defer db.DropTable(&entity.Wallet{}, &entity.WalletAccount{})

	err = godotenv.Load("../../.env") // Update the path accordingly
	require.NoError(t, err)

	// Redis connection options
	options := &rds.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Default DB
	}

	// // initiate redis client for testing
	client := rds.NewClient(options)
	require.NoError(t, err)

	r := repositories.NewRepository(db)
	walletAccountService := service.NewWalletAccountService(r, client)

	err = seeder.SeedWalletAccountFaker(db)
	require.NoError(t, err)

	t.Run("if token is valid and exist, expected no error", func(t *testing.T) {
		var (
			token = "Token 1d54be82cbad3ec463dcd0ceab26fb409d5e4a52"
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletAccountService.EnableWallet(ctx, tokenString)
		require.NoError(t, err)
		require.Equal(t, err, nil)
		assert.NotNil(t, res)
	})

	t.Run("if token is no valid or not exist, expected error", func(t *testing.T) {
		var (
			token = "Token 1d54be82cbad3ec463dcd0ceab26fb409d5e4b57"
		)
		tokenString, err := helpers.ParseTokenHex(token)
		require.NoError(t, err)

		res, err := walletAccountService.EnableWallet(ctx, tokenString)
		require.Error(t, err)
		assert.Nil(t, res)
	})
}
