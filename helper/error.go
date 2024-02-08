package helper

import "log"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalIfError(err error) {
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
