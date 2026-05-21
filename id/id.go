package id

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func New(prefix string) string {
	return prefix + "_" + ulid.Make().String()
}

func NewRaw() string {
	return ulid.Make().String()
}

func MustParse(s string) (ulid.ULID, error) {
	parts := strings.Split(s, "_")
	if len(parts) == 2 {
		return ulid.Parse(parts[1])
	}
	return ulid.Parse(s)
}
