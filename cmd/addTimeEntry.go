package cmd

import (
	"fmt"

	"github.com/bradmccoydev/teamwork-cli/teamwork"
	"github.com/spf13/cobra"
)

// addChildCmd represents the addChildCmd command
var addTimeEntryCmd = &cobra.Command{
	Use:   "ADD_PARTNER",
	Short: "Add a partner",
	Long:  `Add a Partner`,
	Run: func(cmd *cobra.Command, args []string) {
		err := addTimeEntry(args)
		if err != nil {
			fmt.Println(err)
		}
	},
}

var params []string

func init() {
	rootCmd.AddCommand(addTimeEntryCmd)
	addTimeEntryCmd.PersistentFlags().StringArrayVar(&params, "time", params, "Time Parameters")
}

// AddChild - For adding a child
func addTimeEntry(args []string) error {
	personToAddPartner := args[0]
	partnerName := args[1]

	familyTree, err := teamwork.CreateTimeEntryForProject()

	if err != nil {
		fmt.Println(err)
	}

	teamwork.CreateTimeEntryForProject(familyTree, personToAddPartner, partnerName)

	return nil
}
