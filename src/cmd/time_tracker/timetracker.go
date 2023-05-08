package time_tracker

import (
	"fmt"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"timetracker/cmd/time_tracker/flags"
	_authDelivery "timetracker/internal/Auth/delivery"
	authRep "timetracker/internal/Auth/repository/redis"
	authUsecase "timetracker/internal/Auth/usecase"
	_entryDelivery "timetracker/internal/Entry/delivery"
	entryRep "timetracker/internal/Entry/repository/postgres"
	entryUsecase "timetracker/internal/Entry/usecase"
	_goalDelivery "timetracker/internal/Goal/delivery"
	goalRep "timetracker/internal/Goal/repository/postgres"
	goalUsecase "timetracker/internal/Goal/usecase"
	_projectDelivery "timetracker/internal/Project/delivery"
	projectRep "timetracker/internal/Project/repository/postgres"
	projectUsecase "timetracker/internal/Project/usecase"
	_tagDelivery "timetracker/internal/Tag/delivery"
	tagRep "timetracker/internal/Tag/repository/postgres"
	tagUsecase "timetracker/internal/Tag/usecase"
	userRep "timetracker/internal/User/repository/postgres"
	"timetracker/internal/middleware"
)

type TimeTracker struct {
	base
	PostgresClient flags.PostgresFlags `toml:"postgres-client"`
	RedisClient    flags.RedisFlags    `toml:"redis-client"`
	Server         flags.ServerFlags   `toml:"server"`
}

func (tt TimeTracker) Run() error {
	e := echo.New()
	services, err := tt.Init(e)

	if err != nil {
		return fmt.Errorf("can not init services: %w", err)
	}

	postgresClient, err := tt.PostgresClient.Init()

	if err != nil {
		return fmt.Errorf("can not connect to Postgres client: %w", err)
	}

	redisClient, err := tt.RedisClient.Init()

	if err != nil {
		return fmt.Errorf("can not connect to Redis client: %w", err)
	}

	logger := services.Logger

	entryRepo := entryRep.NewEntryRepository(postgresClient)
	userRepo := userRep.NewUserRepository(postgresClient)
	tagRepo := tagRep.NewTagRepository(postgresClient)
	goalRepo := goalRep.NewGoalRepository(postgresClient)
	projectRepo := projectRep.NewProjectRepository(postgresClient)
	authRepo := authRep.NewAuthRepository(redisClient)

	entryUC := entryUsecase.New(entryRepo, tagRepo, userRepo)
	goalUC := goalUsecase.New(goalRepo)
	projectUC := projectUsecase.New(projectRepo)
	tagUC := tagUsecase.New(tagRepo)
	authUC := authUsecase.New(userRepo, authRepo)

	_entryDelivery.NewDelivery(e, entryUC)
	_goalDelivery.NewDelivery(e, goalUC)
	_projectDelivery.NewDelivery(e, projectUC)
	_tagDelivery.NewDelivery(e, tagUC)
	_authDelivery.NewDelivery(e, authUC)

	e.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: `time=${time_custom} remote_ip=${remote_ip} ` +
			`host=${host} method=${method} uri=${uri} user_agent=${user_agent} ` +
			`status=${status} error="${error}" ` +
			`bytes_in=${bytes_in} bytes_out=${bytes_out}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(echoMiddleware.Recover())
	authMiddleware := middleware.NewMiddleware(authUC)
	e.Use(authMiddleware.Auth)

	httpServer := tt.Server.Init(e)
	server := Server{*httpServer}

	if err := server.Start(); err != nil {
		logger.Fatal(err)
	}
	return nil
}
