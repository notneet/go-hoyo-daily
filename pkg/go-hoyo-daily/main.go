package main

import (
	"fmt"
	"os"

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

func doJob() {
	var (
		resp *resty.Response
		err  error

		hoyoCookie    = os.Getenv("HOYOLAB_COOKIES")
		hoyoSignInUrl = os.Getenv("HOYOLAB_SIGN_IN_URL")

		genshinActId = os.Getenv("GENSHIN_ACT_ID")
		honkaiSR     = os.Getenv("HONKAI_SR_ACT_ID")
		zenlessZZ    = os.Getenv("ZZZ_ACT_ID")
		honkaiI3Rd   = os.Getenv("HI_3RD_ACT_ID")
		tot          = os.Getenv("TOT_ACT_ID")
	)

	if genshinActId != "" {
		resp, err = helper.ApiClientHoyo(hoyoSignInUrl, hoyoCookie, genshinActId).Post("")
		helper.PanicIfError(err)

		fmt.Println(resp, "Genshin")
	}

	if honkaiSR != "" {
		resp, err = helper.ApiClientHoyo(hoyoSignInUrl, hoyoCookie, honkaiSR).Post("")
		helper.PanicIfError(err)

		fmt.Println(resp, "Honkai SR")
	}

	if zenlessZZ != "" {
		resp, err = helper.ApiClientHoyo(hoyoSignInUrl, hoyoCookie, zenlessZZ).Post("")
		helper.PanicIfError(err)

		fmt.Println(resp, "ZZZ")
	}

	if honkaiI3Rd != "" {
		resp, err = helper.ApiClientHoyo(hoyoSignInUrl, hoyoCookie, honkaiI3Rd).Post("")
		helper.PanicIfError(err)

		fmt.Println(resp, "HI3")
	}

	if tot != "" {
		resp, err = helper.ApiClientHoyo(hoyoSignInUrl, hoyoCookie, tot).Post("")
		helper.PanicIfError(err)

		fmt.Println(resp, "TOT")
	}
}
