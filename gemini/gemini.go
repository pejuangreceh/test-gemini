package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewClient(question string) (any, error) {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyC4sd-ZAD7CcLRfHhrS3jwMyoStSiZGg8A"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(question))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	data, _ := json.Marshal(resp)
	fmt.Println(resp)

	var jsonData any
	err2 := json.Unmarshal(data, &jsonData)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err2
	}

	return jsonData, nil
}
