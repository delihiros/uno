package main

import (
	"fmt"

	"github.com/delihiros/uno/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("uno")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading config file: %w\n", err))
	}
	cmd.Execute()
}
