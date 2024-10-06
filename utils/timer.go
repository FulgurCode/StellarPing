package utils

import (
	"time"
)

func OnceADay(fn func()) {
	// fn()

	go func() {
		for {
			time.Sleep(time.Hour * 24)

			fn()
		}
	}()
}
