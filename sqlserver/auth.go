package sqlserver

import (
	"database/sql"
	"fmt"
	"github.com/satit13/bcsync/auth"
)

// NewAuthRepository creates new auth repository
func NewAuthRepository(db *sql.DB) (auth.Repository, error) {
	r := authRepository{db}
	return &r, nil
}

type authRepository struct {
	db *sql.DB
}

func (repo *authRepository) GetToken(tokenID string) (*auth.Token, error) {
	var m struct {
		ClientID    sql.NullInt64
		AccountID   sql.NullInt64
		VendingID   sql.NullInt64
		VendingUUID sql.NullString
		TokenID     sql.NullString
	}

	err := repo.db.QueryRow(`
		select
			accounts.client_id, account_id,  null as vending_id, null as vending_uuid,accesses.id
			from accesses inner join accounts on accesses.account_id = accounts.id
			where accesses.id = $1
		union select
			client_id, null, vending_id, vending_uuid, token
		from vending_access
		where token = $1
	`, tokenID).Scan(
		&m.ClientID, &m.AccountID, &m.VendingID, &m.VendingUUID, &m.TokenID,
	)
	if err == sql.ErrNoRows {
		return nil, auth.ErrTokenNotFound
	}
	if err != nil {
		return nil, err
	}

	tk := auth.Token{ID: tokenID}
	fmt.Println("postgres.auth.go -> auth.Token.ID = ", tokenID)
	//if m.ClientID.Valid {
	//	tk.ClientID = m.ClientID.Int64
	//} else {
	//	tk.ClientID = -1
	//}
	if m.AccountID.Valid {
		tk.AccountID = m.AccountID.Int64
	} else {
		tk.AccountID = -1
	}
	//if m.VendingID.Valid {
	//	tk.VendingID = m.VendingID.Int64
	//} else {
	//	tk.VendingID = -1
	//}
	//if m.VendingUUID.Valid {
	//	tk.VendingUUID = m.VendingUUID.String
	//}
	if m.TokenID.Valid {
		tk.TokenID = m.TokenID.String
	}
	fmt.Println("return sqlserver.auth.GetToken")
	return &tk, nil
}
