package database

import (
    "github.com/spf13/viper"
)

// const Login = "root"
// const Passwrd = ""
// const Address = "127.0.0.1:3306"
// const Name = "simple_website"
type DataBaseParams struct {
    Login string
    Passwrd string
    Address string
    Port string
    Name string
}

func (db *DataBaseParams) initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")

    err := viper.ReadInConfig()
    if err != nil {
        return err
    }

    db.Login = viper.GetString("login")
    db.Passwrd = viper.GetString("passwrd")
    db.Address = viper.GetString("address")
    db.Port = viper.GetString("port")
    db.Name = viper.GetString("name")

    return nil
}

func ParseConfig() (*DataBaseParams, error) {
    db := &DataBaseParams{}
    err := db.initConfig()
    if err != nil {
        return nil, err
    }
    return db, nil
}
