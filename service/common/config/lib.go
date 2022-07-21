package config


import (
	"os"
	"encoding/json"
)

func Parsing[T any](dst *T) error {
	if len(os.Args) <1 {
		return os.ErrNotExist
	}
	src := os.Args[1]
	return json.Unmarshal([]byte(src),dst)
}