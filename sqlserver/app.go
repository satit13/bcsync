package sqlserver

import (
	"database/sql"
	"github.com/satit13/bcsync/app"
)

// NewMobileRepository creates new mobile repository
func NewAppeRepository(db *sql.DB) (app.Repository, error) {
	r := appRepo{db}
	return &r, nil
}

type appRepo struct {
	db *sql.DB
}

func (r *appRepo) AddInvoice(inv *app.Invoice) (interface{}, error) {
	return true, nil
}
