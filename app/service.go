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
	resp, err := s.repo.AddInvoice(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
