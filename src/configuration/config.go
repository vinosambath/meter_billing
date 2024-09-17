package configuration

import (
	"MeterBilling/src/constants"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database *Database `yaml:"database"`
	Logger   *Logger   `yaml:"logger"`
}

type Database struct {
	NumberOfDatabases int                    `yaml:"number_of_databases"`
	Connections       []*DatabaseConnections `yaml:"connections"`
}

type Logger struct {
	Path string `yaml:"path"`
}

type DatabaseConnections struct {
	DSN string `yaml:"dsn"`
}

var ConfigInstance Config
var configSingleton sync.Once

func InitConfig() {
	configSingleton.Do(func() {
		config := Config{}
		ConfigInstance = config.parseAndLoadConfig()
	})
}

func GetConfig() Config {
	return ConfigInstance
}

func (c Config) parseAndLoadConfig() Config {
	config := &Config{}
	workingDir, _ := os.Getwd()

	file, err := os.Open(workingDir + constants.CONFIG_FILE_PATH)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		panic(err)
	}

	return *config
}
