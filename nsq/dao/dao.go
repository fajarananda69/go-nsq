package dao

import "log"

func FindError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}
