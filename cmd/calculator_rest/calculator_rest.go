package main

import (
	"Calculator/internal/controller"
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
)

func SetViperConfig() {
	viper.SetEnvPrefix("calcr")
	viper.SetDefault("port", "8080")
	err := viper.BindEnv("port")
	if err != nil {
		slog.Info("Used default port", "port", viper.GetString("CalcPort"))
	}
	viper.AutomaticEnv()
}

func main() {
	SetViperConfig()

	http.HandleFunc("/calc", controller.RestHandler)

	slog.Info("Listening on port " + viper.GetString("port"))
	err := http.ListenAndServe(":"+viper.GetString("port"), nil)
	if err != nil {
		return
	}
}
