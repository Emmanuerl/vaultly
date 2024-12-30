package wallets

import (
	"context"
	"errors"
	"time"

	"github.com/emmanuerl/vaultly/pkg/internal"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/uptrace/bun"
)

// required args when creating a wallet
type CreateArgs struct {
	Currency  string
	AccountID string
}

var ErrExistingWallet = errors.New("a wallet already exists for the given account id")

type Wallet struct {
	bun.BaseModel `bun:"table:wallets"`
	ID            string    `bun:",scanonly" json:"id"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
	AccountID     string    `json:"account_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type Repo struct {
	db *bun.DB
}

func NewRepo(db *bun.DB) *Repo {
	return &Repo{db}
}

// Create records a new wallet entry. It errors on duplicate AccountID detection
func (w *Repo) Create(ctx context.Context, args *CreateArgs) (*Wallet, error) {
	wallet := &Wallet{AccountID: args.AccountID, Currency: args.Currency}

	_, err := w.db.NewInsert().Model(wallet).Returning("*").Exec(ctx)

	if err, ok := err.(*pgconn.PgError); ok && err.Code == internal.DbErrUniqueViolation {
		return nil, ErrExistingWallet
	}

	return wallet, err

}
