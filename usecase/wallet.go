package usecase

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fajarcandraaa/mini_wallet_exercise/helpers"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/dto"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/entity"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/presentation"
	"github.com/fajarcandraaa/mini_wallet_exercise/internal/service"
	"github.com/pkg/errors"
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

func (u *WalletUseCase) EnabledWallet(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("enabledWallet")
		ctx       = context.Background()
		token     = r.Header.Get("Authorization")
	)

	tokenString, err := helpers.ParseTokenHex(token)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, err.Error())
		return
	}

	service, err := u.service.WalletAccount.EnableWallet(ctx, tokenString)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrPermissionNotAllowed:
			responder.FieldErrors(w, err, http.StatusUnauthorized, err.Error())
			return
		case entity.ErrWalletAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusBadRequest, err.Error())
			return
		}
	}

	response := dto.ToResponse("success", service)
	responder.SuccessJSON(w, response, http.StatusCreated, "success")
	return
}

func (u *WalletUseCase) ViewBallance(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("viewBallanceOnWallet")
		ctx       = context.Background()
		token     = r.Header.Get("Authorization")
	)

	tokenString, err := helpers.ParseTokenHex(token)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, err.Error())
		return
	}

	service, err := u.service.WalletAccount.ViewBallanceOnWallet(ctx, tokenString)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrPermissionNotAllowed:
			responder.FieldErrors(w, err, http.StatusUnauthorized, err.Error())
			return
		case entity.ErrWalletAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusBadRequest, err.Error())
			return
		}
	}

	response := dto.ToResponse("success", service)
	responder.SuccessJSON(w, response, http.StatusOK, "success")
	return
}

func (u *WalletUseCase) TopUpBalance(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("topUpBallanceOnWallet")
		ctx       = context.Background()
		token     = r.Header.Get("Authorization")
		amount    = r.FormValue("amount")
		reffId    = r.FormValue("reference_id")
	)

	amountInt, _ := strconv.Atoi(amount)

	tokenString, err := helpers.ParseTokenHex(token)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, err.Error())
		return
	}

	service, err := u.service.WalletTransaction.TopUpVirtualMoney(ctx, amountInt, reffId, tokenString)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrPermissionNotAllowed:
			responder.FieldErrors(w, err, http.StatusUnauthorized, err.Error())
			return
		case entity.ErrWalletAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusBadRequest, err.Error())
			return
		}
	}

	response := dto.ToResponse("success", service)
	responder.SuccessJSON(w, response, http.StatusCreated, "success")
	return
}

func (u *WalletUseCase) WithdrawlBalance(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("withdrawlBallanceOnWallet")
		ctx       = context.Background()
		token     = r.Header.Get("Authorization")
		amount    = r.FormValue("amount")
		reffId    = r.FormValue("reference_id")
	)

	amountInt, _ := strconv.Atoi(amount)

	tokenString, err := helpers.ParseTokenHex(token)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, err.Error())
		return
	}

	service, err := u.service.WalletTransaction.UseVirtualMoney(ctx, amountInt, reffId, tokenString)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrPermissionNotAllowed:
			responder.FieldErrors(w, err, http.StatusUnauthorized, err.Error())
			return
		case entity.ErrWalletAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusBadRequest, err.Error())
			return
		}
	}

	response := dto.ToResponse("success", service)
	responder.SuccessJSON(w, response, http.StatusCreated, "success")
	return
}
