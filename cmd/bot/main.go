package main

import (
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/notneet/go-hoyo-daily/pkg/config"
	"github.com/notneet/go-hoyo-daily/pkg/logger"
	"github.com/notneet/go-hoyo-daily/pkg/sentry"
)

func main() {
	logger := logger.InitLogger()

	err := bootstrap(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type App struct {
	config *config.Config
	// db     *database.DB
	logger *slog.Logger
	wg     sync.WaitGroup
}

func bootstrap(logger *slog.Logger) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// db, err := database.NewDB(cfg.Dsn)
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	sqlDB, err := db.DB.DB()
	// 	if err != nil {
	// 		logger.Error("failed to get sql.DB", "error", err)
	// 	} else {
	// 		sqlDB.Close()
	// 	}
	// }()

	err = sentry.InitSentry(cfg.SentryDsn)
	if err != nil {
		return err
	}

	app := &App{
		config: cfg,
		// db:     db,
		logger: logger,
	}

	return app.setupRunner()
}
