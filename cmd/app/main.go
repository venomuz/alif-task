package main

import (
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/migration"
	"github.com/venomuz/alif-task/internal/service"
	"github.com/venomuz/alif-task/internal/storage/psqlrepo"
	"github.com/venomuz/alif-task/internal/transport/rest"
	"github.com/venomuz/alif-task/pkg/database"
	"github.com/venomuz/alif-task/pkg/hash"
	"github.com/venomuz/alif-task/pkg/logger"
)

func main() {
	// Initialize logger zap
	logger.New("debug", "app")

	// Load configs
	cfg, err := config.Init("configs")
	if err != nil {
		logger.Zap.Fatal("error while load configs", logger.Error(err))
		return
	}

	// Connection to Database Postgresql
	DB, err := database.NewClient(cfg)
	if err != nil {
		logger.Zap.Fatal("error while connect to database", logger.Error(err))
		return
	}

	// Gorm Auto Migration
	err = migration.AutoMigrate(DB)
	if err != nil {
		logger.Zap.Fatal("error while migration tables", logger.Error(err))
		return
	}

	// Initialize Repositories Postgresql
	psqlRepos := psqlrepo.NewRepositories(DB)

	hasher := hash.NewPasswordHasher()

	// Initialize Services
	services := service.NewServices(service.Deps{
		PsqlRepo: psqlRepos,
		Cfg:      cfg,
		Hash:     hasher,
	})

	// Rest API Load Handlers
	handlers := rest.NewHandler(services, cfg)

	// Initialize Rest API handlers
	srv := handlers.Init()

	// Run listener HTTP server
	err = srv.Run(cfg.HTTP.Host + ":" + cfg.HTTP.Port)
	if err != nil {
		logger.Zap.Fatal("error while start server", logger.Error(err))
		return
	}
}
