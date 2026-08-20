package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Order struct {
	ID        int
	VendorID  int
	CourierID int
	ClientID  int
	Address   string
	TimeDate  string //пока string
}

type Client struct {
	ID    int
	Name  string
	Phone string //пока string
}

type Vendor struct {
	ID      int
	Name    string
	Address string
	Phone   string // пока string
}

type CourierID struct {
	ID    int
	Name  string
	Phone string
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, `
		Статус %d 
		сервер работает
		круто`, http.StatusOK)
	})

	http.HandleFunc("/flowers", flowers())

	text, err := os.ReadFile("orders.jso")
	if err != nil {
		log.Println("failed to open file")
	}
	fmt.Println(string(text))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	//1. /health endpoint для проверки

}

func flowers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷
		      💐🌸💮🏵️🌹🥀🌺🌻🌼🌷`,
		)
	}
}
