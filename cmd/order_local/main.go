package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rcasachi/studies-golang/internal/infra/database"
	"github.com/rcasachi/studies-golang/internal/usecases"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	usecase := usecases.NewCalculateFinalPrice(orderRepository)

	input := usecases.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := usecase.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
