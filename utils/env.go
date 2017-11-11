package utils

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func Env(key, valueIfNotFound string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return valueIfNotFound
}

func EnvInt(key string, valueIfNotFound int) int {
	s := Env(key, "")
	if s == "" {
		return valueIfNotFound
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		e := errors.Wrapf(err, "Parse env value failed (%v)", key)
		panic(e)
	}
	return n
}
