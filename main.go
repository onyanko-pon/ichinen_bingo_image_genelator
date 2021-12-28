package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/onyanko-pon/ichinen-bingo_image_genelator/html_builder"
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

	e.GET("/image.png", func(c echo.Context) error {
		var todoList []string = []string{}

		for i := 0; i < 25; i++ {
			todoList = append(todoList, fmt.Sprintf("バンジーをする %d", i))
		}
		data := html_builder.HTMLData{
			ImageURL: "https://pbs.twimg.com/profile_images/1399403028755619841/JqRHZEkb_400x400.jpg",
			TodoList: todoList,
		}
		html := html_builder.BuildHTML(data)
		fmt.Println("html", html)
		imageData, _ := bingoImageGenelator.GenImage(c.Request().Context(), html)
		return c.Stream(http.StatusOK, "image/png", imageData)
	})

	e.Logger.Fatal(
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))),
	)
}
