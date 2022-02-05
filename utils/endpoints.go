package utils

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type EndpointResponse struct {
	Method   string `yaml:"method"`
	HTTPCode int    `yaml:"httpCode"`
	Body     string `yaml:"body"`
}

type Endpoint struct {
	Name        string             `yaml:"name"`
	Path        string             `yaml:"path"`
	ContentType string             `yaml:"contentType"`
	Methods     []string           `yaml:"methods"`
	Responses   []EndpointResponse `yaml:"responses"`
}

func GetEndpoints() (Endpoint, error) {
	endpoint := viper.GetStringMap("endpoints")
	log.Printf("name: %v", endpoint["name"].(string))
	log.Printf("path: %v", endpoint["path"].(string))
	// log.Printf("contentType: %v", endpoint["contentType"].(string))
	methods := endpoint["methods"].([]interface{})
	for _, method := range methods {
		log.Printf("method: %v", method.(string))
	}

	// if !ok {
	return Endpoint{}, errors.New("could not convert config type to endpoint")
	// }
	// return endpoint, nil
}
