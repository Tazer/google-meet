package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// roomsCmd represents the rooms command
var roomsCmd = &cobra.Command{
	Use:   "rooms",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly rooms a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rooms called")
		resourceCall, err := adminService.Resources.Calendars.List(viper.GetString("customer")).Do()

		if err != nil {
			log.Fatalf("Error listing calendars err: %v", err)
		}

		for _, cal := range resourceCall.Items {
			log.Printf("Cal: %+v", cal)
		}
	},
}

func init() {
	rootCmd.AddCommand(roomsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// roomsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// roomsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
