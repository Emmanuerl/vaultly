package api

import (
	"errors"
	"net/http"

	"github.com/emmanuerl/vaultly/pkg/config"
	"github.com/emmanuerl/vaultly/pkg/internal"
	"github.com/emmanuerl/vaultly/pkg/wallets"
	"github.com/go-chi/chi/v5"
	v "github.com/go-ozzo/ozzo-validation"
)

type createWalletDTO struct {
	AccountID string `json:"account_id" validate:"required"`
	Currency  string `json:"currency" validate:"required"`
}

func (c createWalletDTO) Validate() error {
	return v.ValidateStruct(&c,
		v.Field(&c.AccountID, v.Required),
		v.Field(&c.Currency, v.Required),
	)
}

func WalletRoutes(app *config.App) *chi.Mux {
	repo := wallets.NewRepo(app.Db)

	r := chi.NewRouter()
	wc := newWalletcontroller(repo)

	r.Post("/wallets", wc.create)
	return r
}

type walletController struct {
	repo *wallets.Repo
}

func newWalletcontroller(repo *wallets.Repo) *walletController {
	return &walletController{repo: repo}
}

func (wc *walletController) create(w http.ResponseWriter, r *http.Request) {
	var dto createWalletDTO
	internal.ParseAndValidate(r, &dto)

	wallet, err := wc.repo.Create(r.Context(), &wallets.CreateArgs{
		Currency:  dto.Currency,
		AccountID: dto.AccountID,
	})

	if err == nil {
		internal.HttpRespond(w, http.StatusCreated, wallet)
		return
	}

	appErr := &internal.ApiErr{Message: err.Error(), StatusCode: http.StatusInternalServerError}

	if errors.Is(err, wallets.ErrExistingWallet) {
		appErr.StatusCode = http.StatusConflict
	}

	panic(appErr)
}
