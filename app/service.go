package app

// Service is the mobile service
type Service interface {
	AddInvoice(invoice *Invoice) (interface{}, error)
}

// NewService creates new mobile service
func NewService(repo Repository) (Service, error) {
	s := service{repo}
	return &s, nil
}

type service struct {
	repo Repository
}

func (s *service) AddInvoice(req *Invoice) (interface{}, error) {
	s.repo.AddInvoice(req)
	return true, nil
}
