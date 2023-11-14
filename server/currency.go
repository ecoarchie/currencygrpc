package server

import (
	"context"

	protos "github.com/ecoarchie/currencygrpc/protos/currency"

	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l, protos.UnimplementedCurrencyServer{}}
}

func (c *Currency) GetRate(ctx context.Context, preq *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Habgle GetRate", "base", preq.GetBase(), "destination", preq.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
