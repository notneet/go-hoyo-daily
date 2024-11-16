package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/notneet/go-hoyo-daily/helper"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	helper.FatalIfError(err)

	// force run for first time
	doJob()

	c := cron.New()
	c.AddFunc("0 5 * * *", doJob) // run at 5 subuh
	c.Start()

	fmt.Println("service is running!")
	select {}
}

type HoyoCheckInOpt struct {
	Name      string
	SignInURL string
	ActID     string
}

func doJob() {
	var (
		resp *resty.Response
		err  error

		hoyoCookie = os.Getenv("HOYOLAB_COOKIES")
	)

	envVarConfigs := []HoyoCheckInOpt{
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
			resp, err = helper.ApiClientHoyo(hoyoCookie, config.ActID, strings.ToLower(config.Name)).Post(config.SignInURL)
			helper.PanicIfError(err)

			fmt.Println(resp, config.Name)
		}
	}
}
