package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	ClusterID        string `mapstructure:"CLUSTER_ID"`
	ClientID         string `mapstructure:"CLIENT_ID"`
	NatsURL          string `mapstructure:"NATS_URL"`
}

func LoadEnvVariables() *Config {
	var c Config
	//tell viper the path of env file
	viper.AddConfigPath("./")
	//tell viper the name of file
	viper.SetConfigName("app")
	//tell viper type of file
	viper.SetConfigType("env")

	//reads all the variables from env file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Error reading env file", err)
	}

	//unmarshal the loaded env variables
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(&c)
	return &c
}
