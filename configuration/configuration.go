package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/myevents/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
	RestfulTLSEPDefault = "localhost:9191"
)

type ServiceConfig struct {
	DatabaseType        dblayer.DBTYPE `json:"databasetype"`
	DBConnection        string         `json:"dbconnection"`
	RestfulEndpoint     string         `json:"restfulapi_endpoint"`
	RestfultTLSEndpoint string         `json:"restfulapi_tlsendpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DatabaseType:        DBTypeDefault,
		DBConnection:        DBConnectionDefault,
		RestfulEndpoint:     RestfulEPDefault,
		RestfultTLSEndpoint: RestfulTLSEPDefault,
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}
	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
