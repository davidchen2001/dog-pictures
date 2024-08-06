package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"io"
	"encoding/json"
	"bytes"
)

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