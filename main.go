package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/ichinen-bingo_image_genelator/image_genelator"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	bingoImageGenelator := image_genelator.ImageGenalator{
		ImageQuority: 100,
		ImageHeight:  1200,
		ImageWidth:   800,
	}

	e.GET("/image", func(c echo.Context) error {
		imageData, _ := bingoImageGenelator.GenImage(c.Request().Context(), "hoghoge")
		return c.Stream(http.StatusOK, "image/png", imageData)
	})

	e.Logger.Fatal(
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))),
	)
}
