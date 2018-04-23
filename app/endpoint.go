package app

import (
	"context"
	"github.com/gin-gonic/gin"
)

func makeAddInvoice(s Service) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		return gin.H{"message": "endpoint success"}, nil
	}
}
