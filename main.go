package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/EDDYCJY/go-gin-example/routers"
	_ "github.com/go-sql-driver/mysql"
	"tsmc.com/go-gin-docker/repository"
)

func DBConnection() (db *sql.DB) {
	// service_url := "oracle://GA731852:uwygnnr2@192.168.1.114:1521/ORCLCDB.localdomain"
	db, err := sql.Open("mysql", "root:usbw@tcp(192.168.1.105:3306)/classicmodels")
	if err != nil {
		fmt.Println("1err", err)
	}
	return db
}

func main() {

	mariaDB := DBConnection()
	defer mariaDB.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	mariaRepo := repository.NewMariaRepository(mariaDB)
	res, err := mariaRepo.GetCustomers(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		fmt.Println(v)
	}
	// port := os.Getenv("PORT")
	routersInit := routers.InitRouter()
	readTimeout := 60 * time.Second
	writeTimeout := 60 * time.Second
	endPoint := "8000"
	maxHeaderBytes := 1 << 20
}
