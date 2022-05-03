package usecase

import (
	"context"
	"time"

	"tsmc.com/go-gin-docker/domain"
)

type MariaUsecase struct {
	mariaRepo      domain.MariaRepository
	contextTimeout time.Duration
}

func NewMariaUsecase(m domain.MariaRepository, timeout time.Duration) domain.MariaUsecase {
	return &MariaUsecase{
		mariaRepo:      m,
		contextTimeout: timeout,
	}
}

func (m *MariaUsecase) GetCustomers(ctx context.Context) (customers []*domain.Customers, err error) {
	ctx, cancel := context.WithTimeout(ctx, m.contextTimeout)
	defer cancel()
	customers, err = m.mariaRepo.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return
}
