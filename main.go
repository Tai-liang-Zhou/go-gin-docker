package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"tsmc.com/go-gin-docker/Customer/repository"
	"tsmc.com/go-gin-docker/Customer/routers"
	_ "tsmc.com/go-gin-docker/docs"
)

func DBConnection() (db *sql.DB) {
	// service_url := "oracle://GA731852:uwygnnr2@192.168.1.114:1521/ORCLCDB.localdomain"
	db, err := sql.Open("mysql", "root:test@tcp(localhost:3306)/classicmodels")
	if err != nil {
		fmt.Println("1err", err)
	}
	return db
}

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name tlchoud
// @contact.url test

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
func main() {

	mariaDB := DBConnection()
	defer mariaDB.Close()

	mariaRepo := repository.NewMariaRepository(mariaDB)

	r := gin.Default()
	// gin.SetMode(gin.ReleaseMode)
	routers.NewMariaRouter(r, mariaRepo)
	//TODO: add mariadb resource to mac
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
