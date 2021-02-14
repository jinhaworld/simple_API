// Simple package that illustrates basic usage of echo
package main

import (
	"flag"
	"fmt"
	"simple-echo/internal/db"
	"simple-echo/internal/web"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var (
	connectionDSN = flag.String("dsn", "root:test@tcp(library_db:3306)/library_db?parseTime=true", "Database string")
	// connectionDSN = flag.String("dsn", "root:test@tcp(localhost:3306)/products?parseTime=true", "Database string")
)

type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func main() {
	flag.Parse()

	_, err := db.GetDB(*connectionDSN)
	if err != nil {
		log.WithError(err).Error("Can't create db connection")
		return
	}

	// Pointer to echo
	e := echo.New()
	v := validator.New()
	e.Validator = &ProductValidator{validator: v}

	web.SetRoutes(e)

	// Port, handle error
	err = e.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}

	// e.Logger.Print("Listening on port 8080")
	// Log and print the error
	// e.Logger.Fatal(e.Start(":8080"))
}
