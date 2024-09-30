package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	// application "test-gemini"
)

func AskGemini() {
	ctx := context.Background()
	fmt.Println(os.Getenv("API_KEY"))
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyC4sd-ZAD7CcLRfHhrS3jwMyoStSiZGg8A"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("How do you describe Indonesia, answer it in bahasa Indonesia"))
	if err != nil {
		log.Fatal(err)
	}
	data, _ := json.Marshal(resp)
	fmt.Println(resp)

	var jsonData any
	err2 := json.Unmarshal(data, &jsonData)
	if err2 != nil {
		log.Fatal(err2)
	}
	// fmt.Println(data)
	fmt.Println(jsonData)
	// responses.ResponseData
}

func main() {

	ctx := context.Background()
	fmt.Println(os.Getenv("API_KEY"))
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyC4sd-ZAD7CcLRfHhrS3jwMyoStSiZGg8A"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("How do you describe Indonesia, answer it in bahasa Indonesia"))
	if err != nil {
		log.Fatal(err)
	}
	data, _ := json.Marshal(resp)
	fmt.Println(resp)

	var jsonData any
	err2 := json.Unmarshal(data, &jsonData)
	if err2 != nil {
		log.Fatal(err2)
	}
	// fmt.Println(data)
	fmt.Println(jsonData)

	e := echo.New()
	// v1 := "api/elloy/v1"

	// e.POST(v1 + "tanya-gemini-wkwk")
	e.Logger.Fatal(e.Start("0.0.0.0:" + "8001"))

}
