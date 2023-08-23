package service

import (
	"context"
	"regexp"

	"github.com/fajarcandraaa/mini_wallet_exercise/internal/dto"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/repositories"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type walletService struct {
	repo *repositories.Repository
}

func NewWalletService(repo *repositories.Repository) *walletService {
	return &walletService{
		repo: repo,
	}
}

var _ WalletServiceContract = &walletService{}

// CreateAccount implements WalletServiceContract.
func (s *walletService) CreateAccount(ctx context.Context, payload presentation.InitiateWalletAccountRequest) (*presentation.InitiateWalletAccountResponse, error) {
	err := validation.ValidateStruct(&payload,
		validation.Field(&payload.CustomerXid, validation.Required, validation.Match(regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)), validation.Length(3, 100)),
	)
	if err != nil {
		return nil, err
	}
	
	payloadData := dto.CustomerXidToDatabase(payload)
	result, err := s.repo.Wallet.StoreNewWallet(ctx, *payloadData)
	if err != nil {
		return nil, err
	}

	resp := dto.TokenToResponse(*result)
	return &resp, nil
}
