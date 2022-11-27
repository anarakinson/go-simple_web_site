package database

import (
    "os"
    "github.com/spf13/viper"
)


type DataBaseParams struct {
    Login string
    Passwrd string
    Address string
    Port string
    Name string
}

func (db *DataBaseParams) initConfig() error {
    db.Login = viper.GetString("database.login")
    db.Passwrd = os.Getenv("MYSQL_PASSWORD")
    db.Address = viper.GetString("database.address")
    db.Port = viper.GetString("database.port")
    db.Name = viper.GetString("database.name")

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
