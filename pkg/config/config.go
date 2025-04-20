package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Http_server struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production

type Config struct {
	Env          string `yaml:"env" env-required:"true"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	Http_server  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}

	_, err := os.Stat(configPath)

	if os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	err = cleanenv.ReadConfig(configPath, &cfg)


	if err != nil {
		log.Fatalf("can not read config file : %s", err.Error())
	}

	return &cfg
}
