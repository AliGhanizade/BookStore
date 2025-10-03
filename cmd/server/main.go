// cmd/server/main.go
package main

import (
	"BookStore/config"
	"BookStore/internal/domain"
	"BookStore/internal/repository"
	"BookStore/internal/service"
	"BookStore/internal/transport"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	
	db, err := gorm.Open(postgres.Open(config.LoadDBConfig().DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	if err := db.AutoMigrate(&domain.User{}, &domain.Book{}); err != nil {
		log.Fatal("failed to migrate:", err)
	}


	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)

	r := gin.Default()

	userHandler := transport.NewUserHandler(userService)
	userHandler.RegisterRoutes(r)

	bookHandler := transport.NewBookHandler(bookService)
	bookHandler.RegisterRoutes(r)

	port := 8080
	fmt.Printf("Server running on :%d\n", port)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal("failed to run server:", err)
	}
}