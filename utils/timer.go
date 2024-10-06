package utils

import (
	"time"
)

func OnceADay(fn func(string)) {
	fn("7")

	go func() {
		for {
			time.Sleep(time.Hour * 24)

			fn("1")
		}
	}()
}
