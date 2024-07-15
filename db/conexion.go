package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string = "host=localhost user=postgres password=maglio dbname=apigolang port=5432"
var Database *gorm.DB

func Iniciar_conexion() {
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("CONEXION EXITOSA")
	}
}

func Cerrar_conexion() {
	if Database != nil {
		sw, err := Database.DB()
		if err != nil {
			panic(err)
		} else {
			sw.Close()
			fmt.Println("CONEXION CERRADA")
		}
	}
}
