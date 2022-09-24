package repository

import (
	"context"
	"database/sql"
	"fmt"

	"tsmc.com/go-gin-docker/domain"
)

type mariaRepo struct {
	DB *sql.DB
}

func NewMariaRepository(db *sql.DB) domain.MariaRepository {
	return &mariaRepo{
		DB: db,
	}
}

func (m *mariaRepo) GetCustomers(ctx context.Context) (customer []*domain.Customers, err error) {
	query := `SELECT customerName,phone,addressLine1 FROM classicmodels.customers;`
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	resData := make([]*domain.Customers, 0)

	for rows.Next() {
		rowData := domain.Customers{}
		err := rows.Scan(
			&rowData.Name,
			&rowData.Phone,
			&rowData.Address,
			&rowData.CreateTime,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(rowData)
		resData = append(resData, &rowData)
	}

	return resData, nil

}
