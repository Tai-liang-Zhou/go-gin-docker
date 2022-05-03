package usecase

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"tsmc.com/go-gin-docker/Customer/repository"
	"tsmc.com/go-gin-docker/domain"
)

func TestMariaUsecase_GetCustomers(t *testing.T) {
	type fields struct {
		mariaRepo      domain.MariaRepository
		contextTimeout time.Duration
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		beforTest     func(sqlmock.Sqlmock)
		wantCustomers []*domain.Customers
		wantErr       bool
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
			wantCustomers: []*domain.Customers{{Name: sql.NullString{String: "tlchoud", Valid: true},
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
			mariaRepo := repository.NewMariaRepository(db)
			m := &MariaUsecase{
				mariaRepo:      mariaRepo,
				contextTimeout: time.Minute * 15,
			}
			if tt.beforTest != nil {
				tt.beforTest(mockSQL)
			}

			gotCustomers, err := m.GetCustomers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MariaUsecase.GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomers, tt.wantCustomers) {
				t.Errorf("MariaUsecase.GetCustomers() = %v, want %v", gotCustomers, tt.wantCustomers)
			}
		})
	}
}
