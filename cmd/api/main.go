package main

import (
	"encoding/json"
	"net/http"

	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rcasachi/studies-golang/internal/entities"
)

func main() {
	// // chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Order)

	// http.HandleFunc("/order", Order)
	// http.ListenAndServe(":8888", r)

	e := echo.New()
	e.GET("/order", OrderForEcho)
	// http.ListenAndServe(":8888", e)
	e.Logger.Fatal(e.Start(":8888"))
}

func OrderForEcho(c echo.Context) error {
	order, err := entities.NewOrder("1234", 1000, 1)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func Order(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Method not Allowed")
		return
	}

	// w.Write([]byte("Hello World"))
	order, err := entities.NewOrder("1234", 1000, 1)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
