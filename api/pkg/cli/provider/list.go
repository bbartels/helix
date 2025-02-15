package provider

import (
	"fmt"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/helixml/helix/api/pkg/client"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List provider endpoints",
	Long:    ``,
	RunE: func(cmd *cobra.Command, _ []string) error {
		apiClient, err := client.NewClientFromEnv()
		if err != nil {
			return err
		}

		endpoints, err := apiClient.ListProviderEndpoints(cmd.Context())
		if err != nil {
			return fmt.Errorf("failed to list provider endpoints: %w", err)
		}

		table := tablewriter.NewWriter(cmd.OutOrStdout())

		header := []string{"ID", "Name", "Description", "Type", "Owner", "Base URL", "Created"}

		table.SetHeader(header)

		table.SetAutoWrapText(false)
		table.SetAutoFormatHeaders(true)
		table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetCenterSeparator("")
		table.SetColumnSeparator("")
		table.SetRowSeparator("")
		table.SetHeaderLine(false)
		table.SetBorder(false)
		table.SetTablePadding(" ")
		table.SetNoWhiteSpace(false)

		for _, e := range endpoints {
			row := []string{
				e.ID,
				e.Name,
				e.Description,
				string(e.EndpointType),
				fmt.Sprintf("%s (%s)", e.Owner, e.OwnerType),
				e.BaseURL,
				e.Created.Format(time.RFC3339),
			}

			table.Append(row)
		}

		table.Render()

		return nil
	},
}
