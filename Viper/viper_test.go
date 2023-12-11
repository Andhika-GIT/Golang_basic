package viper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	config := viper.New()

	// read the config.json
	config.SetConfigName("Config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	// get the value
	appName := config.GetString("app.name")
	fmt.Println(appName)
}

func TestENV(t *testing.T) {
	config := viper.New()

	// read the env
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)

	// get the value
	appName := config.GetString("APP_NAME")
	fmt.Println(appName)
}
