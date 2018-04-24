package mssql

type BcArInvoice struct {
	DocNO   string `db:"docno"`
	DocDate string `db:"docdate"`
	ArCode  string `db:"arcode"`
	ArName  string `json:"ar_name"`
}
