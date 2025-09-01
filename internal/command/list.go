package command

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

type RatesResponse struct {
	Quotes map[string]float64 `json:"quotes"`
}

var base string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List exchange rates from a base currency",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("https://api.exchangerate.host/live?source=%s&access_key=%s", base, apiKey)
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		var res RatesResponse
		if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
			return err
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"Currency", "Rate"})
		for k, v := range res.Quotes {
			table.Append([]string{k, fmt.Sprintf("%.4f", v)})
		}
		table.Render()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&base, "base", "b", "USD", "base currency")
}
