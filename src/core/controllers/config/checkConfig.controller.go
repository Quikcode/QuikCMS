package config

import (
	"QuikCMS/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Config Structure control data configuration
type Config struct {
	Name     string `json:"name" ` // Organization name
	Database struct {
		Database string `json:"db"`
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Email struct {
		Server   string `json:"server"`
		Port     int    `json:"port"`
		Mail     string `json:"mail"`
		Password string `json:"password"`
	} `json:"email"`
}

func CheckConfig(path *utils.Files) bool {
	file := path.Read()
	defer file.Close()

	// Datastore JSON
	var config []Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		if err == io.EOF {
			return false
		}
		panic(err)
	}
	fmt.Println(config)
	return true
}

func CheckConfigTest(path *utils.Files) bool {
	config := []Config{{
		Name: "Org1",
		Database: struct {
			Database string `json:"db"`
			Host     string `json:"host"`
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Database: "database1",
			Host:     "host1",
			Username: "username1",
			Password: "password1",
		},
		Email: struct {
			Server   string `json:"server"`
			Port     int    `json:"port"`
			Mail     string `json:"mail"`
			Password string `json:"password"`
		}{
			Server:   "server1",
			Port:     443,
			Mail:     "mail@org.com",
			Password: "password1",
		},
	}}
	configJSON, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}
	// Crea un archivo para escribir el JSON
	file, err := os.Create(path.Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Escribe el JSON en el archivo
	_, err = file.Write(configJSON)
	if err != nil {
		panic(err)
	}
	return false
}
