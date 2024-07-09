package main

import (
	"context"
	"log"

	"video-rec/adapters/gemini"
	"video-rec/config/variables"
)

func main() {
	envVars, err := variables.GetEnvVariables()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := gemini.NewClient(envVars.GeminiAPIKey, ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	uploadedFile, err := gemini.UploadFile(envVars.VideoPath, client, ctx)
	if err != nil {
		log.Fatal(err)
	}

	model := gemini.NewModel(client, envVars.Model)

	promptResponses, err := gemini.PromptVideo(uploadedFile.URI, model, envVars.Prompt, ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Response: %s\n", *promptResponses[0].Content)

	log.Println("Done!")

}
