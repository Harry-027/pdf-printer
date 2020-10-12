package utils

import "log"

// Error check utils ...

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func LogErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
