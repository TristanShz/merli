package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func TranslateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "translate",
		Short: "Translate a file to a target language",
		Run: func(cmd *cobra.Command, args []string) {
			authKey, ok := os.LookupEnv("DEEPL_API_KEY")

			if !ok {
				fmt.Println("DEEPL_API_KEY environment variable is not set")
				return
			}

			fmt.Printf("Translate command %v", authKey)
		},
	}

	return cmd
}
