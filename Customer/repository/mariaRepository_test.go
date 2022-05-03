package repository

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"tsmc.com/go-gin-docker/domain"
)

func Test_mariaRepo_GetCustomers(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		beforTest    func(sqlmock.Sqlmock)
		wantCustomer []*domain.Customers
		wantErr      bool
	}{
		{name: "response normal",
			args: args{context.TODO()},
			beforTest: func(mockSQL sqlmock.Sqlmock) {
				rows := mockSQL.
					NewRows([]string{"customerName", "phone", "addressLine1"}).
					AddRow("tlchoud", "0912345678", "kaohsiung")
				prep := mockSQL.ExpectPrepare("SELECT customerName,phone,addressLine1 FROM classicmodels.customers")
				prep.ExpectQuery().WillReturnRows(rows)

			},
			wantCustomer: []*domain.Customers{{Name: sql.NullString{String: "tlchoud", Valid: true},
				Phone:   sql.NullString{String: "0912345678", Valid: true},
				Address: sql.NullString{String: "kaohsiung", Valid: true}}},
			wantErr: false},
		{name: "response error",
			args: args{context.TODO()},
			beforTest: func(mockSQL sqlmock.Sqlmock) {
				prep := mockSQL.ExpectPrepare("SELECT customerName,phone,addressLine1 FROM classicmodels.customers")
				prep.ExpectQuery().WillReturnError(errors.New("reponse err"))

			},
			wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mockSQL, _ := sqlmock.New()
			defer db.Close()
			m := &mariaRepo{
				DB: db,
			}
			if tt.beforTest != nil {
				tt.beforTest(mockSQL)
			}

			gotCustomer, err := m.GetCustomers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("mariaRepo.GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("mariaRepo.GetCustomers() = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func TestNewMariaRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want domain.MariaRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMariaRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMariaRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
