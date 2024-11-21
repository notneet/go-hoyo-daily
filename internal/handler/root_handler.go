package handler

import (
	"context"
	"log/slog"

	"github.com/getsentry/sentry-go"
	"github.com/notneet/go-hoyo-daily/internal/service"
	"github.com/robfig/cron/v3"
)

type RootHandlerInterface interface {
	Run(ctx context.Context) error
	Stop(ctx context.Context) error
}

type RootHandler struct {
	logger    *slog.Logger
	sentryHub *sentry.Hub
	service   service.RootServiceInterface
}

func NewRootHandler(logger *slog.Logger, service service.RootServiceInterface, sentryHub *sentry.Hub) RootHandlerInterface {
	return &RootHandler{
		service:   service,
		logger:    logger,
		sentryHub: sentryHub,
	}
}

func (h *RootHandler) Run(ctx context.Context) error {
	h.logger.Info("bot started running")

	c := cron.New(cron.WithSeconds()) // Enable second-level precision

	// force run
	h.logger.Info("running first task")
	err := h.service.ProcessCheckIn(ctx)
	if err != nil {
		h.logger.Error("failed to run first task", "error", err)
		h.sentryHub.CaptureException(err)

		return err
	}

	// Schedule the task to run at 5 AM every day
	_, err = c.AddFunc("0 0 5 * * *", func() {
		h.logger.Info("executing task at 5 AM")

		err := h.service.ProcessCheckIn(ctx)
		if err != nil {
			h.logger.Error("failed to execute task", "error", err)
		}
	})
	if err != nil {
		h.logger.Error("failed to schedule task", "error", err)
		return err
	}

	// Start the cron scheduler
	c.Start()
	defer c.Stop()

	// Wait until the goroutine returns
	<-ctx.Done()
	return ctx.Err()
}

func (h *RootHandler) Stop(ctx context.Context) error {
	select {
	case <-ctx.Done():
		h.logger.Error("bot stopping due to context deadline", "error", ctx.Err())
		return ctx.Err()
	default:
		h.logger.Info("bot stopped successfully")
		return nil
	}
}
