package main

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var config = struct {
	AppName string `default:"app name"`
	AppURL  string

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func main() {
	configor.Load(&config, "config.yml")
	fmt.Printf("config: %v\n", config)
}
