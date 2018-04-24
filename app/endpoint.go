package app

import (
	"context"
)


func makeAddInvoice(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		_invoice := &Invoice{}
		resp, err := s.AddInvoice(_invoice)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
