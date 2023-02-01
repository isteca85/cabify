package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	application "github.com/isteca85/car-pooling-challenge/pkg/application/http"
	"github.com/isteca85/car-pooling-challenge/pkg/infrastructure"
	"log"
	"net/http"
)

func main() {
	log.Println("Opening database...")
	b := &infrastructure.Bbdd{}
	var err = b.OpenDdbb()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	defer b.CloseDdbb()
	log.Println("Cleaning database...")
	//b.CleanDdbb()

	log.Println("Init server...")
	s := &application.Server{}
	s.InitServer()
	s.DataBase = &*b

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
