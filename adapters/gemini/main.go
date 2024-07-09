package gemini

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func NewClient(apiKey string, ctx context.Context) (*genai.Client, error) {
	return genai.NewClient(ctx, option.WithAPIKey(apiKey))
}

func UploadFile(filePath string, client *genai.Client, ctx context.Context) (*genai.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	opts := genai.UploadFileOptions{DisplayName: ""}
	response, err := client.UploadFile(ctx, "", f, &opts)
	if err != nil {
		log.Fatal(err)
	}

	var file *genai.File = response
	fmt.Printf("Uploaded file %s as: %q\n", file.DisplayName, file.URI)

	response, err = client.GetFile(ctx, file.Name)
	if err != nil {
		log.Fatal(err)
	}

	for response.State == genai.FileStateProcessing {
		fmt.Print(".")
		time.Sleep(10 * time.Second)

		response, err = client.GetFile(ctx, file.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println()

	fmt.Printf("File %s is ready for inference as: %q\n", response.DisplayName, response.URI)

	return response, nil
}

func NewModel(client *genai.Client, modelName string) *genai.GenerativeModel {
	return client.GenerativeModel(modelName)
}

func PromptVideo(videoURI string, model *genai.GenerativeModel, prompt string, ctx context.Context) ([]*genai.Candidate, error) {
	constructedPrompt := []genai.Part{
		genai.FileData{URI: videoURI},
		genai.Text(prompt),
	}

	resp, err := model.GenerateContent(ctx, constructedPrompt...)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Candidates, nil
}
