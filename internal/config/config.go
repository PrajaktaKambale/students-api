package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"  env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"production" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func Load() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path tp the configuration file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("missing config path")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("failed to read config: %s", err)
	}
	return &cfg

}
