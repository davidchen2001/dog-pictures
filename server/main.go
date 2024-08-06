package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"io"
)

type DogImage struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

const API_URL = "https://dog.ceo/api/breeds/image/random"

func main() {
	app := fiber.New()

	app.Get("/api/dog", getImages)

	log.Fatal(app.Listen(":4000"))
	fmt.Println("Server running")
	
}

func getImages(c *fiber.Ctx) error {
	resp, err := http.Get(API_URL)
	if err != nil {
		log.Fatalln(err)
	 }

	 defer resp.Body.Close()
	 body, err := io.ReadAll(resp.Body)
	 fmt.Println(resp)

	 return c.JSON(body)
}