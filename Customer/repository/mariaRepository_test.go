package repository

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gopkg.in/guregu/null.v4"
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
					NewRows([]string{"customerName", "phone", "addressLine1", "create_time"}).
					AddRow("tlchoud", "0912345678", "kaohsiung", time.Date(2022, 9, 24, 12, 0, 0, 0, time.UTC))
				prep := mockSQL.ExpectPrepare("SELECT customerName,phone,addressLine1 FROM classicmodels.customers")
				prep.ExpectQuery().WillReturnRows(rows)

			},
			wantCustomer: []*domain.Customers{
				{
					Name:       null.StringFrom("tlchoud").NullString,
					Phone:      null.StringFrom("0912345678").NullString,
					Address:    null.StringFrom("kaohsiung").NullString,
					CreateTime: null.TimeFrom(time.Date(2022, 9, 24, 12, 0, 0, 0, time.UTC)).NullTime,
				},
			},
			wantErr: false},
		{name: "response error",
			args: args{context.TODO()},
			beforTest: func(mockSQL sqlmock.Sqlmock) {
				prep := mockSQL.ExpectPrepare("SELECT customerName,phone,addressLine1 FROM classicmodels.customers")
				prep.ExpectQuery().WillReturnError(errors.New("reponse err"))

			},
			wantErr: true},
		{name: "No any row response",
			args: args{context.TODO()},
			beforTest: func(mockSQL sqlmock.Sqlmock) {
				prep := mockSQL.ExpectPrepare("SELECT customerName,phone,addressLine1 FROM classicmodels.customers")
				prep.ExpectQuery().WillReturnError(sql.ErrNoRows)

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
