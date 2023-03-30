package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/migration"
	"github.com/venomuz/alif-task/internal/repository/psqlrepo"
	"github.com/venomuz/alif-task/internal/repository/redisrepo"
	"github.com/venomuz/alif-task/internal/service"
	"github.com/venomuz/alif-task/internal/transport/rest"
	"github.com/venomuz/alif-task/internal/transport/rest/server"
	"github.com/venomuz/alif-task/pkg/auth"
	"github.com/venomuz/alif-task/pkg/database"
	"github.com/venomuz/alif-task/pkg/hash"
	"github.com/venomuz/alif-task/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// main
//	@title						Alif task
//	@version					1.0
//	@description				This is a sample server app server.
//	@contact.name				API Support
//	@contact.url				https://t.me/xalmatoff
//	@contact.email				venom.uzz@mail.ru
//	@securityDefinitions.apikey	BearerAuth
//	@Description				Authorization for accounts
//	@in							header
//	@name						Authorization
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

	rdbC := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS.Host + fmt.Sprintf(":%d", cfg.REDIS.Port),
		Password: cfg.REDIS.Password, // no password set
	})

	// Gorm Auto Migration
	err = migration.AutoMigrate(DB)
	if err != nil {
		logger.Zap.Fatal("error while migration tables", logger.Error(err))
		return
	}

	// Initialize Repositories Postgresql
	psqlRepos := psqlrepo.NewRepositories(DB)

	redsRepos := redisrepo.NewRedisRepo(rdbC)

	// Initialize hasher
	hasher := hash.NewPasswordHasher()

	// Initialize Token Manager
	tokenManager := auth.NewTokenManager(cfg.AUTH.JwtSigningKey)

	// Initialize Services
	services := service.NewServices(service.Deps{
		PsqlRepo:     psqlRepos,
		RedisRepo:    redsRepos,
		Cfg:          cfg,
		Hash:         hasher,
		TokenManager: tokenManager,
	})

	// Rest API Load Handlers
	handlers := rest.NewHandler(services, cfg)

	// Initialize Rest API handlers
	srv := server.NewServer(cfg, handlers.Init())

	// Run server via goroutine
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Zap.Error("error occurred while running http server", logger.Error(err))
		}
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			panic(err)
		}
		time.Local = loc
	}()

	logger.Zap.Info("server is started")

	// Graceful Shutdown
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Zap.Info("shutdown server")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	//stop server
	if err := srv.Stop(ctx); err != nil {
		logger.Zap.Fatal("filed to stop server", logger.Error(err))
	}
}
