package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"robot-app/internal/http/handler"
	"robot-app/internal/http/server"
	"robot-app/internal/repository"
	"robot-app/pkg/config"
	"robot-app/pkg/database"
	"robot-app/validation"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Server",
	Run: func(cmd *cobra.Command, args []string) {

		s := setupServer()
		if err := s.Start(); err != nil {
			log.Printf("Failed to start HTTP server with error: %s\n", err.Error())
		}
	},
}

func setupServer() server.Server {
	db, err := database.New(config.GetDatabaseConfig().DbUrl)
	if err != nil {
		log.Printf("Failed database.New: %s\n", err.Error())
		panic(err)
	}

	database.Migrate(db)
	database.Seed(db)

	validate := validation.New()

	validate.RegisterValidate()

	repo := repository.NewRepository(db)

	return server.New(handler.New(validate, repo))
}
