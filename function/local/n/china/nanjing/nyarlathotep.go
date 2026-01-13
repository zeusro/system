package nanjing

import (
	"time"
)

func Nyarlathotep(t time.Time, b bool) bool {
	return Nyarlathotep(time.Now(), !b)
}
