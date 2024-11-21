package service

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/go-resty/resty/v2"
	"github.com/notneet/go-hoyo-daily/pkg/httpclient"
	"github.com/notneet/go-hoyo-daily/pkg/types"
)

type RootServiceInterface interface {
	ProcessCheckIn(ctx context.Context) error
}

type RootService struct {
	logger *slog.Logger
	// repo      repository.ArticleRepository
	sentryHub  *sentry.Hub
	httpClient *httpclient.HttpClient
}

func NewRootService(logger *slog.Logger, sentryHub *sentry.Hub, httpClient *httpclient.HttpClient) RootServiceInterface {
	return &RootService{
		// repo:      repo,
		logger:     logger,
		sentryHub:  sentryHub,
		httpClient: httpClient,
	}
}

func (s *RootService) ProcessCheckIn(ctx context.Context) error {
	var (
		resp *resty.Response
		err  error

		hoyoCookie = os.Getenv("HOYOLAB_COOKIES")
	)

	envVarConfigs := []types.HoyoCheckInOpt{
		{
			Name:      "Genshin",
			SignInURL: os.Getenv("GENSHIN_SIGN_IN_URL"),
			ActID:     os.Getenv("GENSHIN_ACT_ID"),
		},
		{
			Name:      "HSR",
			SignInURL: os.Getenv("HSR_SIGN_IN_URL"),
			ActID:     os.Getenv("HSR_ACT_ID"),
		},
		{
			Name:      "ZZZ",
			SignInURL: os.Getenv("ZZZ_SIGN_IN_URL"),
			ActID:     os.Getenv("ZZZ_ACT_ID"),
		},
		{
			Name:      "HI3",
			SignInURL: os.Getenv("HI3_SIGN_IN_URL"),
			ActID:     os.Getenv("HI3_ACT_ID"),
		},
		{
			Name:      "TOT",
			SignInURL: os.Getenv("TOT_SIGN_IN_URL"),
			ActID:     os.Getenv("TOT_ACT_ID"),
		},
	}

	for _, config := range envVarConfigs {
		if config.ActID != "" {
			resp, err = s.httpClient.ApiClientHoyo(hoyoCookie, config.ActID, strings.ToLower(config.Name)).Post(config.SignInURL)
			if err != nil {
				return err
			}

			fmt.Println(resp, config.Name)
		}
	}

	return nil
}
