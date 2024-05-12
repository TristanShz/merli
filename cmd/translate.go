package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tristanshz/merli/internal/deepl"
)

func getFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func TranslateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "translate [file] [target-lang]",
		Short: "Translate a file to a target language",
		RunE: func(cmd *cobra.Command, args []string) error {
			authKey, ok := os.LookupEnv("DEEPL_API_KEY")

			if !ok {
				fmt.Println("DEEPL_API_KEY environment variable is not set")
				return nil
			}

			filePath := args[0]
			targetLang := args[1]

			content, err := getFileContent(filePath)
			if err != nil {
				fmt.Printf("Error reading file content %v", err)
				return nil
			}

			translator := deepl.NewTranslator(authKey, "free")

			result, err := translator.Translate([]string{content}, targetLang)
			if err != nil {
				fmt.Printf("Error translating content %v", err)
				return nil
			}

			fmt.Println(result.Translations)

			return nil
		},
	}

	return cmd
}
