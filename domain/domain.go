package domain

import (
	"context"
	"database/sql"
)

type Customers struct {
	Name    sql.NullString
	Phone   sql.NullString
	Address sql.NullString
}

type MariaRepository interface {
	GetCustomers(ctx context.Context) (customer []*Customers, err error)
}
type MariaUsecase interface {
	GetCustomers(ctx context.Context) (customer []*Customers, err error)
}
