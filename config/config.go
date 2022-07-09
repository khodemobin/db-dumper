package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Storage struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Path     string
	Driver   string
}

type Database struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Driver   string
}

type Task struct {
	ArchiveDriver   string   `yaml:"archiveDriver"`
	ArchivePassword string   `yaml:"archivePassword"`
	Storages        []string `yaml:"storages"`
	Database        string   `yaml:"database"`
}

type Config struct {
	Storages  []Storage
	Databases []Database
	Tasks     []Task
}

var cfg *Config

func LoadConfig(filePath string) (*Config, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	data = replaceVariablesWithEnv(data)

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

func replaceVariablesWithEnv(data []byte) []byte {
	// todo ability to replace variables from env

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}
	//
	//log.Fatalln(string(data))
	//re := regexp.MustCompile("a(?P<1W>x*)b")

	return data
}
