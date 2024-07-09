package variables

import (
	"os"

	"github.com/joho/godotenv"
)

type Variables struct {
	VideoPath    string
	Prompt       string
	Model        string
	GeminiAPIKey string
}

func GetEnvVariables() (*Variables, error) {
	err := godotenv.Load(".env")

	variables := &Variables{
		VideoPath:    os.Getenv("VIDEO_PATH"),
		Prompt:       os.Getenv("PROMPT"),
		Model:        os.Getenv("MODEL"),
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
	}

	return variables, err
}
