package usecase

import (
	"context"
	"net/http"

	"github.com/fajarcandraaa/mini_wallet_exercise/helpers"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/dto"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
)

type WalletUseCase struct {
	service *service.Service
}

func NewWalletUseCase(service *service.Service) *WalletUseCase {
	return &WalletUseCase{
		service: service,
	}
}

func (u *WalletUseCase) InitializeAccountWallet(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("initializeAccountWallet")
		ctx       = context.Background()
		param     = r.FormValue("customer_xid")
	)

	payload := &presentation.InitiateWalletAccountRequest{
		CustomerXid: param,
	}
	walletService, err := u.service.WalletService.CreateAccount(ctx, *payload)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.ToResponse("success", walletService)
	responder.SuccessJSON(w, response, http.StatusCreated, "success")
	return
}
