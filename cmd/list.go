package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly list a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		calendarID := "primary"
		fmt.Println("list called")

		t := time.Now().Format(time.RFC3339)
		events, err := calendarService.Events.List(calendarID).ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(1).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		}
		//fmt.Println("Upcoming events:")
		if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
		} else {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				// json, err := item.MarshalJSON()

				// if err != nil {
				// 	log.Fatalf("Marshal not owrking: %v", err)
				// }

				fmt.Printf("Your adhoc google meet link: %s\n", item.ConferenceData.EntryPoints[0].Uri)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
