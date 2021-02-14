package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

// Validation struct
type Body struct {
	Name string `json:"name" validate:"required,min=4"`
	// Name string `json:"name" validate:"required,min=4,email"`
	// Vendor string `json:"vendor" validate:"min=4,max=10"`
	// // If something isprovided with vendor, email should also be provided
	// Email           string `json:"email" validate:"required_with=Vendor,email"`
	// Website         string `json:"website" validate:"url"`
	// Country         string `json:"country" validate:"len=2"`
	// DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
}

var products = map[int]string{
	1: "car",
	2: "computer",
	3: "phone",
}

// very simple hello handler with string
// Try with POST and see that it doesn't work
func HelloHandler(c echo.Context) error {
	return c.String(200, "HELLO")
	// return c.String(http.StatusOK, "HELLO")
}

func GetProductsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func GetProductByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	product, exists := products[id]
	if exists {
		return c.JSON(http.StatusOK, product)
	}

	return c.JSON(http.StatusNotFound, fmt.Sprintf("Product for id: %d does not exist", id))
}

func AddProductHandler(c echo.Context) error {
	// go get github.com/go-playground/validator/v10
	// returns validate, a struct for validation
	v := validator.New()

	var reqBody Body
	// Binds the request body with provided type
	err := c.Bind(&reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Validator not registered
	// Register at top
	err = c.Validate(reqBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Validate after you do the binding
	err = v.Struct(reqBody)
	if err != nil {
		// Same as
		// return err
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	products[len(products)+1] = reqBody.Name
	return c.JSON(http.StatusOK, products)
}

// go get github.com/go-playground/validator/v10

func UpdateProductHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var reqBody Body

	if err = c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err = c.Validate(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, exists := products[id]
	if exists {
		products[id] = reqBody.Name
		return c.JSON(http.StatusOK, products)
	}
	return c.JSON(http.StatusBadRequest, "Product does not exist")
}

func DeleteProductHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	_, exists := products[id]
	if exists {
		delete(products, id)
		return c.JSON(http.StatusOK, products)
	}
	return c.JSON(http.StatusBadRequest, "Product does not exist")
}
