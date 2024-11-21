package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/notneet/go-hoyo-daily/internal/handler"
	"github.com/notneet/go-hoyo-daily/internal/service"
	"github.com/notneet/go-hoyo-daily/pkg/httpclient"
)

const (
	defaultShutdownPeriod = 30 * time.Second
)

func (app *App) setupRunner() error {
	shutdownErrorChan := make(chan error)
	ctx := sentry.SetHubOnContext(context.Background(), sentry.CurrentHub().Clone())
	hub := sentry.GetHubFromContext(ctx)

	httpClient := httpclient.NewHttpClient()
	// repo := repository.NewArticleRepository(app.logger, app.db)
	service := service.NewRootService(app.logger, hub, httpClient)
	runner := handler.NewRootHandler(app.logger, service, hub)

	go func() {
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan

		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)
		defer cancel()

		shutdownErrorChan <- runner.Stop(ctx)
	}()

	app.logger.Info("starting bot")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := runner.Run(ctx)
	if err != nil {
		return err
	}

	err = <-shutdownErrorChan
	if err != nil {
		return err
	}

	app.logger.Info("stopped bot")
	app.wg.Wait()

	return nil
}
