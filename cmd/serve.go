package cmd

import (
	"github.com/diptomondal007/go_clean_arch/config"
	"github.com/diptomondal007/go_clean_arch/server"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd command to start the service
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start go clean service",
	Long:  "Start go clean service",

	Run: func(cmd *cobra.Command, args []string) {
		app := server.NewApp(config.GetApp())
		err := app.Run()
		if err != nil {
			log.Fatalf("%s", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}