package database

import (
	"github.com/estrategiaconcursos/UsersTest/config"
	"github.com/spf13/viper"
	"log"
	"os"
	"upper.io/db.v3"
	"upper.io/db.v3/postgresql"
)

// Sess variavel que faz a conexao com o bando de dados
var Sess db.Database

func init() {
	var err error

	viper.SetConfigName("config")
	viper.AddConfigPath(os.Getenv("GOPATH") + "/src/github.com/estrategiaconcursos/UsersTest/config")
	//viper.AddConfigPath(".")

	var configuration config.Configuration

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	var settings = postgresql.ConnectionURL{
		Host:     configuration.Database.Host,
		Database: configuration.Database.Name,
		User:     configuration.Database.User,
		Password: configuration.Database.Pass,
	}

	Sess, err = postgresql.Open(settings)
	if err != nil {
		log.Fatal(err.Error())
	}
}
