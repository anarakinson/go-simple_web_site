package main

import (
    "log"

    "internal/server"

    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
)


func initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}


func main() {
    // parse configs
    err := initConfig()
    if err != nil {
        log.Fatal("[!] Error when parsing configs: %s", err.Error())
    }
    // parse variables
    err = godotenv.Load()
    if err != nil {
        log.Fatal("[!] Error when parsing environment variables: %s", err.Error())
    }

    // start up website
    server.RunServer()
}
