package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"robot-app/pkg/config"
)

var rootCmd = &cobra.Command{
	Short: "robot app",
	Long:  `Robot App`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	config.Load()

	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
