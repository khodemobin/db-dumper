package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Storage struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
	Path     string
	Driver   string
}

type Database struct {
	Name     string
	Host     string
	Port     string
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

	log.Fatalln(cfg)

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

func replaceVariablesWithEnv(data []byte) []byte {
	dataString := string(data)
	if err := godotenv.Load(); err != nil {
		return data
	}

	re := regexp.MustCompile(`\$\{([A-Z]|[a-z]|_)+\w}`)
	for _, item := range re.FindAllString(dataString, -1) {
		envName := findEnvName(item)
		env := os.Getenv(envName)
		dataString = strings.ReplaceAll(dataString, item, env)
	}

	return []byte(dataString)
}

func findEnvName(item string) string {
	var re = regexp.MustCompile(`^\${(\w+)}$`)
	match := re.FindStringSubmatch(item)
	return match[1]
}
