package main

import (
	"os"

	"github.com/rutuj-waghare/currency-converter-cli/internal/command"
)

func main() {
	apiKey := os.Getenv("CURRENCY_API_KEY")
	command.Execute(apiKey)
}
