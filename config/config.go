package config

import (
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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

type Compress struct {
	Driver   string
	Password string
}

type Task struct {
	Compress Compress
	Storages []string `yaml:"storages"`
	Database string   `yaml:"database"`
}

type Config struct {
	TempPath  string
	Storages  []Storage
	Databases []Database
	Tasks     []Task
}

var cfg *Config

func LoadConfig(configPath, envPath string) (*Config, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	data = replaceVariablesWithEnv(data, envPath)

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	cfg.TempPath = ".tmp_dumps"

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

func replaceVariablesWithEnv(data []byte, envPath string) []byte {
	dataString := string(data)
	if err := godotenv.Load(envPath); err != nil {
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
