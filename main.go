package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"io"
	"encoding/json"
	"bytes"
	"os"
	"github.com/joho/godotenv"
	"cors"
)

const API_URL = "https://dog.ceo/api/breeds/image/random"

func main() {

	if os.Getenv("ENV") != "production" {
		// Load the .env file if not in production
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3006",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	app.Get("/api/dog", getImages)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

	fmt.Println("Server running")
	
}

func getImages(c *fiber.Ctx) error {
	resp, err := http.Get(API_URL)
	if err != nil {
		log.Fatalln(err)
	 }

	 defer resp.Body.Close()

	 body, err := io.ReadAll(resp.Body)
	 formattedData := formatJSON(body)

	 fmt.Println(formattedData)

	 return c.JSON(formattedData)
}

func formatJSON(data []byte) string {
    var out bytes.Buffer
    err := json.Indent(&out, data, "", " ")

    if err != nil {
        fmt.Println(err)
    }

    d := out.Bytes()
    return string(d)
}