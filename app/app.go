package app

type Repository interface {
	AddInvoice(invoice *Invoice) (interface{}, error)
}


