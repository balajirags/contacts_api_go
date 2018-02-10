package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
	"regexp"
)

func getIntOrPanic(key string) int {
	checkKey(key)
	v, err := strconv.Atoi(fatalGetString(key))
	panicIfErrorForKey(err, key)
	return v
}

func fatalGetString(key string) string {
	checkKey(key)
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func getString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func getFeature(key string) bool {
	v, err := strconv.ParseBool(fatalGetString(key))
	if err != nil {
		return false
	}
	return v
}

func checkKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("%s key is not set", key)
	}
}

func panicIfErrorForKey(err error, key string) {
	if err != nil {
		log.Fatalf("Could not parse key: %s, Error: %s", key, err)
	}
}

func getBoolWithDefault(key string, defaultVal bool) bool {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}

	boolVal, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal
	}
	return boolVal
}

func GetResponseCodeStub(responseCode int) string {
	str := strconv.Itoa(responseCode)
	reg, err := regexp.Compile(`\d{2}$`)
	if err != nil {
		log.Fatal(err)
	}
	responseCodeStub := reg.ReplaceAllString(str, "xx")
	return responseCodeStub
}
