package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/calendar/v3"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		calendarID := "primary"
		start := time.Now().UTC()

		end := start.Add(10 * time.Minute)
		event := &calendar.Event{

			Summary:     "Adhoc Google Meet",
			Location:    "",
			Description: "Doing an adhoc meeting",
			ConferenceData: &calendar.ConferenceData{
				Notes: "This is a adhoc meeting",
				CreateRequest: &calendar.CreateConferenceRequest{
					RequestId: "1234",
					ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
						Type: "hangoutsMeet",
					},
				},
			},
			Start: &calendar.EventDateTime{
				DateTime: start.Format(time.RFC3339),
			},
			End: &calendar.EventDateTime{
				DateTime: end.Format(time.RFC3339),
			},
		}

		attendees := viper.GetStringSlice("attendees")

		if len(attendees) > 0 {
			event.Attendees = []*calendar.EventAttendee{}

			for _, email := range attendees {
				attendee := &calendar.EventAttendee{
					Email: email,
				}
				event.Attendees = append(event.Attendees, attendee)
			}
		}

		event, err := calendarService.Events.Insert(calendarID, event).ConferenceDataVersion(1).Do()
		if err != nil {
			log.Fatalf("Unable to create event. %v\n", err)
		}

		log.Printf("Created Adhoc meet link conf data: %+v", event.ConferenceData.EntryPoints[0].Uri)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
