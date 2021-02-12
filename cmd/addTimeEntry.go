package cmd

import (
	"fmt"

	"github.com/bradmccoydev/teamwork-cli/model"
	"github.com/bradmccoydev/teamwork-cli/teamwork"
	"github.com/spf13/cobra"
)

var (
	conn *teamwork.Connection
)

// addTimeEntryCmd represents the addTimeEntryCmd command
var addTimeEntryCmd = &cobra.Command{
	Use:   "addtime",
	Short: "Add Time Entry",
	Long:  `Add Time Entry`,
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
	addTimeEntryCmd.PersistentFlags().StringArrayVar(&params, "addtime", params, "Time Parameters")
}

// addTimeEntry - For adding a time entry
func addTimeEntry(args []string) error {
	baseURL := args[0]
	apiToken := args[1]
	projectID := args[2]
	blockOne := args[3]
	blockTwo := args[4]

	conn, err := teamwork.Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
	}

	var timeEntry model.CreateTimeEntryOps
	timeEntry.Description = "0.5 - Daily Standup \n" + blockOne + "\n" + blockTwo
	timeEntry.PersonID = "241410"
	timeEntry.Date = "20210119"
	timeEntry.Time = "09:00:00"
	timeEntry.Hours = "8.5"
	timeEntry.Minutes = "0"
	timeEntry.IsBillable = "false"
	timeEntry.TaskID = ""

	response, err := teamwork.CreateTimeEntryForProject(conn, projectID, &timeEntry)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response.Status)
	fmt.Println(response.ID)

	return nil
}
