package mssql

import (
	"github.com/satit13/bcsync/app"
	"database/sql"
	"fmt"
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
	var arcode string
	rs := r.db.QueryRow("select top 1 code from bcar ")
	rs.Scan(&arcode)
	fmt.Println("arcode ", arcode)
	return map[string]interface{}{
		"ar_code":arcode}, nil
}
