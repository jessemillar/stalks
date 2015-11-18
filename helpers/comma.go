package helpers

import "github.com/dustin/go-humanize"

func Comma(number int) string {
	return humanize.Comma(int64(number))
}
